# Generate a syso file - https://github.com/getlantern/systray#windows
# This is needed as systray uses lxn/walk which requires using a manifest file
rsrc -manifest .\EACS.manifest -ico .\internal\icon\icon.ico -o .\cmd\EACS\EACS.syso

# Have to target cmd\EACS and not cmd\EACS\main.go so `go build` sees the syso file
go build -ldflags "-s -w -H=windowsgui" -o .\EACS.exe -v .\cmd\EACS
