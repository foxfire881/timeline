set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go build -ldflags="-s -w -buildid= -extldflags=-Wl,--gc-sections" -trimpath -gcflags="all=-l"  -o "bin\timeline" main.go