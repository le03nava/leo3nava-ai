param(
    [Parameter(Mandatory = $true)]
    [string]$ReportPath,

    [Parameter(Mandatory = $false)]
    [string]$OutputPath,

    [Parameter(Mandatory = $false)]
    [switch]$IncludeFindings
)

$ErrorActionPreference = 'Stop'

function Resolve-ExistingFilePath {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Path
    )

    $resolvedPath = Resolve-Path -LiteralPath $Path -ErrorAction Stop


    if (-not (Test-Path -LiteralPath $resolvedPath.Path -PathType Leaf)) {
        throw "File not found: $Path"
    }

    return $resolvedPath.Path
}

function Resolve-OutputFilePath {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Path
    )

    if ([System.IO.Path]::IsPathRooted($Path)) {
        return $Path
    }

    return Join-Path -Path (Get-Location).Path -ChildPath $Path
}

function Assert-ImportExcelModule {
    if (-not (Get-Module -ListAvailable -Name ImportExcel)) {
        throw "Required PowerShell module 'ImportExcel' was not found. Install it with: Install-Module ImportExcel -Scope CurrentUser"
    }

    Import-Module ImportExcel -ErrorAction Stop
}

function Convert-MarkdownCell {
    param(
        [Parameter(Mandatory = $false)]
        [AllowNull()]
        [string]$Value
    )

    if ($null -eq $Value) {
        return ''
    }

    $cleanValue = $Value.Trim()
    $cleanValue = $cleanValue -replace '<br\s*/?>', "`n"
    $cleanValue = $cleanValue -replace '&lt;', '<'
    $cleanValue = $cleanValue -replace '&gt;', '>'
    $cleanValue = $cleanValue -replace '&amp;', '&'
    $cleanValue = $cleanValue -replace '\\`', '`'
    $cleanValue = $cleanValue -replace '\\\|', '|'
    $cleanValue = $cleanValue -replace '^`([^`]+)`$', '$1'

    return $cleanValue
}

function Split-MarkdownTableLine {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Line
    )

    $trimmedLine = $Line.Trim()

    if ($trimmedLine.StartsWith('|')) {
        $trimmedLine = $trimmedLine.Substring(1)
    }

    if ($trimmedLine.EndsWith('|')) {
        $trimmedLine = $trimmedLine.Substring(0, $trimmedLine.Length - 1)
    }

    $placeholder = [char]0xE000
    $protectedLine = $trimmedLine -replace '\\\|', $placeholder

    return @($protectedLine -split '\|' | ForEach-Object {
        Convert-MarkdownCell -Value ($_ -replace [regex]::Escape([string]$placeholder), '\|')
    })
}

function Convert-MarkdownHeaderToPropertyName {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Header,

        [Parameter(Mandatory = $true)]
        [hashtable]$UsedNames
    )

    $name = Convert-MarkdownCell -Value $Header

    $name = $name -replace '[^\p{L}\p{Nd}]+', ' '
    $name = ($name.Trim() -replace '\s+', ' ')

    if ([string]::IsNullOrWhiteSpace($name)) {
        $name = 'Column'
    }

    $baseName = $name
    $suffix = 2

    while ($UsedNames.ContainsKey($name)) {
        $name = "$baseName $suffix"
        $suffix++
    }

    $UsedNames[$name] = $true
    return $name
}

function Convert-MarkdownTableToObjects {
    param(
        [Parameter(Mandatory = $true)]
        [string[]]$TableLines
    )

    if ($TableLines.Count -lt 3) {
        throw 'Markdown table must contain header, separator, and at least one data row.'
    }

    $usedNames = @{}
    $headers = @(Split-MarkdownTableLine -Line $TableLines[0] | ForEach-Object {
        Convert-MarkdownHeaderToPropertyName -Header $_ -UsedNames $usedNames
    })

    $rows = @()

    foreach ($line in $TableLines[2..($TableLines.Count - 1)]) {
        $cells = @(Split-MarkdownTableLine -Line $line)

        if ($cells.Count -eq 1 -and [string]::IsNullOrWhiteSpace($cells[0])) {
            continue
        }

        $row = [ordered]@{}

        for ($index = 0; $index -lt $headers.Count; $index++) {
            $value = ''

            if ($index -lt $cells.Count) {
                $value = $cells[$index]
            }

            $row[$headers[$index]] = $value
        }

        $rows += [pscustomobject]$row
    }

    return $rows
}

function Get-MarkdownTableAfterHeading {
    param(
        [Parameter(Mandatory = $true)]
        [string[]]$Lines,

        [Parameter(Mandatory = $true)]
        [string[]]$HeadingAliases
    )

    $headingPattern = '^##\s+(' + (($HeadingAliases | ForEach-Object { [regex]::Escape($_) }) -join '|') + ')\s*$'
    $headingIndex = -1

    for ($index = 0; $index -lt $Lines.Count; $index++) {
        if ($Lines[$index] -match $headingPattern) {
            $headingIndex = $index
            break
        }
    }

    if ($headingIndex -lt 0) {
        throw "Could not find heading: $($HeadingAliases -join ' / ')"
    }

    $tableStart = -1


    for ($index = $headingIndex + 1; $index -lt $Lines.Count; $index++) {
        if ($Lines[$index] -match '^#{1,6}\s+') {
            break
        }

        if ($Lines[$index].Trim().StartsWith('|')) {
            $tableStart = $index
            break
        }
    }

    if ($tableStart -lt 0) {
        throw "Could not find table under heading: $($HeadingAliases -join ' / ')"
    }

    $tableLines = @()

    for ($index = $tableStart; $index -lt $Lines.Count; $index++) {
        if (-not $Lines[$index].Trim().StartsWith('|')) {
            break
        }

        $tableLines += $Lines[$index]
    }

    return $tableLines
}

function Get-FencedYamlMetadata {
    param(
        [Parameter(Mandatory = $true)]
        [string[]]$Lines
    )

    $metadata = [ordered]@{}
    $fenceStart = -1

    for ($index = 0; $index -lt $Lines.Count; $index++) {
        if ($Lines[$index] -match '^```yaml\s*$') {
            $fenceStart = $index
            break
        }
    }

    if ($fenceStart -lt 0) {
        return $metadata
    }

    for ($index = $fenceStart + 1; $index -lt $Lines.Count; $index++) {
        if ($Lines[$index] -match '^```\s*$') {
            break
        }

        if ($Lines[$index] -match '^\s*([^:#][^:]*):\s*(.*?)\s*$') {
            $metadata[$matches[1].Trim()] = $matches[2].Trim()
        }
    }

    return $metadata
}

function Convert-MetadataToRows {
    param(
        [Parameter(Mandatory = $true)]
        [System.Collections.IDictionary]$Metadata
    )

    return @($Metadata.GetEnumerator() | ForEach-Object {
        [pscustomobject]@{
            Field = $_.Key
            Value = $_.Value
        }
    })
}

function Get-FindingRows {
    param(
        [Parameter(Mandatory = $true)]
        [object[]]$Rows,

        [Parameter(Mandatory = $true)]
        [string]$SourceTable
    )

    return @($Rows | Where-Object {
        $properties = $_.PSObject.Properties.Name
        $properties -contains 'Finding' -and
        -not [string]::IsNullOrWhiteSpace($_.Finding) -and
        $_.Finding -notmatch '^(none|None|N/A)$'
    } | ForEach-Object {
        $_ | Add-Member -NotePropertyName SourceTable -NotePropertyValue $SourceTable -Force -PassThru
    })
}

$reportFile = Resolve-ExistingFilePath -Path $ReportPath
$reportDirectory = Split-Path -Parent $reportFile
$reportBaseName = [System.IO.Path]::GetFileNameWithoutExtension($reportFile)

if ([string]::IsNullOrWhiteSpace($OutputPath)) {
    $OutputPath = Join-Path -Path $reportDirectory -ChildPath "$reportBaseName.xlsx"
}

$outputFile = Resolve-OutputFilePath -Path $OutputPath
$outputDirectory = Split-Path -Parent $outputFile

if (-not (Test-Path -LiteralPath $outputDirectory -PathType Container)) {
    throw "Output directory not found: $outputDirectory"
}

Assert-ImportExcelModule

$content = Get-Content -LiteralPath $reportFile -Raw -Encoding UTF8
$lines = @($content -split "\r?\n")

$metadata = Get-FencedYamlMetadata -Lines $lines
$securityTableLines = Get-MarkdownTableAfterHeading -Lines $lines -HeadingAliases @('Security Row Validation', 'Compact Security Row Validation')
$sourceTableLines = Get-MarkdownTableAfterHeading -Lines $lines -HeadingAliases @('Corporate Source Row Validation', 'Corporate Source Row Validation Matrix')

$metadataRows = Convert-MetadataToRows -Metadata $metadata
$securityRows = @(Convert-MarkdownTableToObjects -TableLines $securityTableLines)
$sourceRows = @(Convert-MarkdownTableToObjects -TableLines $sourceTableLines)

if ($metadata.Contains('sourceRowExpectedCount')) {
    $expectedSourceRowCount = [int]$metadata['sourceRowExpectedCount']

    if ($sourceRows.Count -ne $expectedSourceRowCount) {
        throw "Source row count mismatch. Expected $expectedSourceRowCount rows, found $($sourceRows.Count)."
    }
}

if (Test-Path -LiteralPath $outputFile -PathType Leaf) {
    Remove-Item -LiteralPath $outputFile -Force
}

$metadataRows | Export-Excel -Path $outputFile -WorksheetName 'Metadata' -AutoSize -FreezeTopRow -BoldTopRow -TableName 'Metadata'
$securityRows | Export-Excel -Path $outputFile -WorksheetName 'Security Rows' -AutoSize -FreezeTopRow -BoldTopRow -TableName 'SecurityRows'
$sourceRows | Export-Excel -Path $outputFile -WorksheetName 'Source Rows' -AutoSize -FreezeTopRow -BoldTopRow -TableName 'SourceRows'

if ($IncludeFindings) {
    $findingRows = @()
    $findingRows += Get-FindingRows -Rows $securityRows -SourceTable 'Security Rows'
    $findingRows += Get-FindingRows -Rows $sourceRows -SourceTable 'Source Rows'

    if ($findingRows.Count -eq 0) {
        $findingRows = @([pscustomobject]@{
            SourceTable = ''
            Finding = 'None'
            Notes = 'No non-empty findings were found.'
        })
    }

    $findingRows | Export-Excel -Path $outputFile -WorksheetName 'Findings' -AutoSize -FreezeTopRow -BoldTopRow -TableName 'Findings'
}

"XLSX created: $outputFile"
"Security rows exported: $($securityRows.Count)"
"Source rows exported: $($sourceRows.Count)"
