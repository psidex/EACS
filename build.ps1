# Builds EACS.exe and packages everything into EACS.zip

Write-Output "`nBuilding .exe"
go build -ldflags "-s -w -H=windowsgui" -o .\EACS.exe -v .\cmd\EACS\main.go

Write-Output "`nPackaging into .zip"
7z a EACS.zip .\EACS.exe .\config-files .\GUIDE.md

Write-Output "`nCleaning up"
Remove-Item .\EACS.exe

Write-Output "`nDone"
