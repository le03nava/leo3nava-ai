$ErrorActionPreference = 'Stop'

if ([string]::IsNullOrWhiteSpace($env:USERPROFILE)) {
    throw 'USERPROFILE is not set. Set USERPROFILE to a valid Windows user profile path before syncing skills.'
}

$repoRoot = Split-Path -Parent $PSScriptRoot
$sourceSkillsPath = Join-Path -Path $repoRoot -ChildPath 'src\skills'

if (-not (Test-Path -LiteralPath $sourceSkillsPath -PathType Container)) {
    throw "Source skills directory not found: $sourceSkillsPath"
}

$opencodeDirectory = Join-Path -Path $env:USERPROFILE -ChildPath '.config\opencode'
$copilotDirectory = 'C:\Users\leo3n\.copilot'

$targets = @(
    @{
        Name = 'opencode'
        RequiredDirectory = $opencodeDirectory
        DestinationPath = Join-Path -Path $opencodeDirectory -ChildPath 'skills'
    },
    @{
        Name = 'copilot'
        RequiredDirectory = $copilotDirectory
        DestinationPath = Join-Path -Path $copilotDirectory -ChildPath 'skills'
    }
)

$syncedTargets = 0

foreach ($target in $targets) {
    if (-not (Test-Path -LiteralPath $target.RequiredDirectory -PathType Container)) {
        "Skipped $($target.Name) skills sync. Required directory not found: $($target.RequiredDirectory)"
        continue
    }

    New-Item -ItemType Directory -Path $target.DestinationPath -Force | Out-Null
    Get-ChildItem -LiteralPath $sourceSkillsPath -Force | Copy-Item -Destination $target.DestinationPath -Recurse -Force

    $syncedTargets++
    "Synced $($target.Name) skills."
    "Source: $sourceSkillsPath"
    "Destination: $($target.DestinationPath)"
}

if ($syncedTargets -eq 0) {
    'No skills were synced because no supported skills directories were found.'
}
