.PHONY:
.SILENT:

filename = "server"

build:
#	 Linux
	GOOS=linux GOARCH=amd64 CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=1 go build -ldflags "-s -H linux" -o ./bin/$(filename)-l64 cmd/server/main.go

#	Windows
#	GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" CGO_CFLAGS="-g -O2 -w" CGO_ENABLED=0 go build -ldflags "-s -H windows" -o ./bin/$(filename)-w64.exe cmd/server/main.go

run: build
	./bin/$(filename)-l64
