$ErrorActionPreference = 'Stop'

if ([string]::IsNullOrWhiteSpace($env:USERPROFILE)) {
    throw 'USERPROFILE is not set. Set USERPROFILE to a valid Windows user profile path before syncing agents.'
}

$repoRoot = Split-Path -Parent $PSScriptRoot
$sourceAgentsPath = Join-Path -Path $repoRoot -ChildPath 'src\agents'

if (-not (Test-Path -LiteralPath $sourceAgentsPath -PathType Container)) {
    throw "Source agents directory not found: $sourceAgentsPath"
}

$opencodeDirectory = Join-Path -Path $env:USERPROFILE -ChildPath '.config\opencode'
$copilotDirectory = 'C:\Users\leo3n\.copilot'

$targets = @(
    @{
        Name = 'opencode'
        RequiredDirectory = $opencodeDirectory
        DestinationPath = Join-Path -Path $opencodeDirectory -ChildPath 'prompts'
    },
    @{
        Name = 'copilot'
        RequiredDirectory = $copilotDirectory
        DestinationPath = Join-Path -Path $copilotDirectory -ChildPath 'agents'
    }
)

$syncedTargets = 0

foreach ($target in $targets) {
    if (-not (Test-Path -LiteralPath $target.RequiredDirectory -PathType Container)) {
        "Skipped $($target.Name) agents sync. Required directory not found: $($target.RequiredDirectory)"
        continue
    }

    New-Item -ItemType Directory -Path $target.DestinationPath -Force | Out-Null
    Get-ChildItem -LiteralPath $sourceAgentsPath -Force | Copy-Item -Destination $target.DestinationPath -Recurse -Force

    $syncedTargets++
    "Synced $($target.Name) agents."
    "Source: $sourceAgentsPath"
    "Destination: $($target.DestinationPath)"
}

if ($syncedTargets -eq 0) {
    'No agents were synced because no supported agent directories were found.'
}
