set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w -buildid= -extldflags=-Wl,--gc-sections" -trimpath -gcflags="all=-l"  -o "bin\timeline.exe" main.go