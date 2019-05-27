build-windows:
	GO111MODULE=on GOOS=windows GOARCH=386 go1.12 build -ldflags -H=windowsgui

build-darwin:
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go1.12 build
