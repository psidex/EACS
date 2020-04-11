# Generate a syso file
# See https://github.com/getlantern/systray#windows for the manifest file
rsrc -manifest EACS.manifest -ico icon\icon.ico -o EACS.syso
# Build - automatically uses *.syso files
# "-s -w" removes debugging stuff
go build -ldflags "-s -w -H=windowsgui" -v .