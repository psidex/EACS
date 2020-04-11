# See https://github.com/getlantern/systray#windows for the manifest file
# Generate a syso file
rsrc -manifest EACS.manifest -ico icon\icon.ico -o EACS.syso
# Build - automatically uses *.syso files
go build -ldflags -H=windowsgui -v .