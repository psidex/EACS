# Generate a syso file - https://github.com/getlantern/systray#windows
# This is needed as systray started using lxn/walk somewhere, which requires using a manifest file
rsrc -manifest EACS.manifest -ico icon\icon.ico -o EACS.syso

# Build (automatically includes *.syso files in the current directory)
# "-s -w" removes debugging stuff
go build -ldflags "-s -w -H=windowsgui" -v .
