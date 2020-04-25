# Builds EACS.exe and packages everything into EACS.zip

# Generate a syso file - https://github.com/getlantern/systray#windows
# This is needed as systray uses lxn/walk which requires using a manifest file
Write-Output "`nGenerating .syso file"
rsrc -manifest .\EACS.manifest -ico .\internal\icon\iconactive.ico -o .\cmd\EACS\EACS.syso

# Have to target cmd\EACS and not cmd\EACS\main.go so `go build` sees the syso file
Write-Output "`nBuilding .exe"
go build -ldflags "-s -w -H=windowsgui" -o .\EACS.exe -v .\cmd\EACS

Write-Output "`nPackaging into .zip"
7z a EACS.zip .\EACS.exe .\config-files .\GUIDE.md

Write-Output "`nCleaning up"
Remove-Item .\EACS.exe

Write-Output "`nDone"
