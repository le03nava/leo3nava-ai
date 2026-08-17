param(
    [Parameter(Mandatory = $true)]
    [string]$InputPath,

    [Parameter(Mandatory = $false)]
    [string]$OutputPath,

    [Parameter(Mandatory = $false)]
    [ValidateSet('png', 'svg', 'pdf')]
    [string]$ImageFormat = 'png',

    [Parameter(Mandatory = $false)]
    [ValidateSet('default', 'neutral', 'dark', 'forest')]
    [string]$Theme = 'neutral',

    [Parameter(Mandatory = $false)]
    [string]$ReferenceDocPath,

    [Parameter(Mandatory = $false)]
    [string]$TableStyle,

    [Parameter(Mandatory = $false)]
    [switch]$KeepTemp,

    [Parameter(Mandatory = $false)]
    [string]$MermaidCliPackage = '@mermaid-js/mermaid-cli'
)

$ErrorActionPreference = 'Stop'

function Resolve-ExistingFilePath {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Path
    )

    $resolvedPath = Resolve-Path -LiteralPath $Path -ErrorAction Stop

    if (-not (Test-Path -LiteralPath $resolvedPath.Path -PathType Leaf)) {
        throw "Input Markdown file not found: $Path"
    }

    return $resolvedPath.Path
}

function Assert-CommandExists {
    param(
        [Parameter(Mandatory = $true)]
        [string]$CommandName,

        [Parameter(Mandatory = $true)]
        [string]$InstallHint
    )

    if (-not (Get-Command -Name $CommandName -ErrorAction SilentlyContinue)) {
        throw "Required command '$CommandName' was not found. $InstallHint"
    }
}

$inputFile = Resolve-ExistingFilePath -Path $InputPath
$inputDirectory = Split-Path -Parent $inputFile
$inputBaseName = [System.IO.Path]::GetFileNameWithoutExtension($inputFile)

$referenceDocFile = $null
if (-not [string]::IsNullOrWhiteSpace($ReferenceDocPath)) {
    $referenceDocFile = Resolve-ExistingFilePath -Path $ReferenceDocPath
}

if ([string]::IsNullOrWhiteSpace($OutputPath)) {
    $OutputPath = Join-Path -Path $inputDirectory -ChildPath "$inputBaseName.docx"
}

if (-not [System.IO.Path]::IsPathRooted($OutputPath)) {
    $OutputPath = Join-Path -Path (Get-Location).Path -ChildPath $OutputPath
}

$outputDirectory = Split-Path -Parent $OutputPath
if (-not (Test-Path -LiteralPath $outputDirectory -PathType Container)) {
    throw "Output directory not found: $outputDirectory"
}

Assert-CommandExists -CommandName 'pandoc' -InstallHint 'Install Pandoc and make sure it is available in PATH.'
Assert-CommandExists -CommandName 'npx' -InstallHint 'Install Node.js/npm so npx is available in PATH.'

$tempRoot = Join-Path -Path $env:TEMP -ChildPath ("md-docx-mermaid-{0}" -f ([Guid]::NewGuid().ToString('N')))
$diagramsDirectory = Join-Path -Path $tempRoot -ChildPath 'diagrams'
$processedMarkdownPath = Join-Path -Path $tempRoot -ChildPath ([System.IO.Path]::GetFileName($inputFile))
$tableStyleFilterPath = Join-Path -Path $tempRoot -ChildPath 'table-style.lua'

New-Item -ItemType Directory -Path $diagramsDirectory -Force | Out-Null

try {
    $markdown = Get-Content -LiteralPath $inputFile -Raw -Encoding UTF8
    $diagramIndex = 0
    $mermaidBlockPattern = '```mermaid\s*\r?\n(?<code>[\s\S]*?)\r?\n```'

    $processedMarkdown = [regex]::Replace(
        $markdown,
        $mermaidBlockPattern,
        {
            param($match)

            $script:diagramIndex++
            $diagramName = 'diagram-{0:D3}' -f $script:diagramIndex
            $diagramSourcePath = Join-Path -Path $diagramsDirectory -ChildPath "$diagramName.mmd"
            $diagramImagePath = Join-Path -Path $diagramsDirectory -ChildPath "$diagramName.$ImageFormat"
            $diagramRelativePath = "diagrams/$diagramName.$ImageFormat"

            Set-Content -LiteralPath $diagramSourcePath -Value $match.Groups['code'].Value -Encoding UTF8

            $arguments = @(
                '-p', $MermaidCliPackage,
                'mmdc',
                '-i', $diagramSourcePath,
                '-o', $diagramImagePath,
                '-t', $Theme
            )

            & npx @arguments

            if ($LASTEXITCODE -ne 0) {
                throw "Mermaid CLI failed while rendering $diagramSourcePath"
            }

            if (-not (Test-Path -LiteralPath $diagramImagePath -PathType Leaf)) {
                throw "Mermaid CLI did not create the expected image: $diagramImagePath"
            }

            return "![Diagram $($script:diagramIndex)]($diagramRelativePath)"
        }
    )

    Set-Content -LiteralPath $processedMarkdownPath -Value $processedMarkdown -Encoding UTF8

    if (-not [string]::IsNullOrWhiteSpace($TableStyle)) {
        $escapedTableStyle = $TableStyle.Replace('\\', '\\\\').Replace('"', '\"')
        $tableStyleFilter = @"
function Table(table)
  table.attributes['custom-style'] = "$escapedTableStyle"
  return table
end
"@
        Set-Content -LiteralPath $tableStyleFilterPath -Value $tableStyleFilter -Encoding UTF8
    }

    $resourcePath = @($inputDirectory, $tempRoot) -join [System.IO.Path]::PathSeparator
    $pandocArguments = @(
        $processedMarkdownPath,
        '-o', $OutputPath,
        "--resource-path=$resourcePath"
    )

    if ($referenceDocFile) {
        $pandocArguments += "--reference-doc=$referenceDocFile"
    }

    if (-not [string]::IsNullOrWhiteSpace($TableStyle)) {
        $pandocArguments += "--lua-filter=$tableStyleFilterPath"
    }

    & pandoc @pandocArguments
    if ($LASTEXITCODE -ne 0) {
        throw 'Pandoc failed while creating the DOCX file.'
    }

    if (-not (Test-Path -LiteralPath $OutputPath -PathType Leaf)) {
        throw "Pandoc did not create the expected DOCX file: $OutputPath"
    }

    "DOCX created: $OutputPath"
    "Mermaid diagrams rendered: $diagramIndex"
    if ($KeepTemp) {
        "Temporary files kept: $tempRoot"
    }
}
finally {
    if (-not $KeepTemp -and (Test-Path -LiteralPath $tempRoot -PathType Container)) {
        Remove-Item -LiteralPath $tempRoot -Recurse -Force
    }
}
