build-windows:
	GO111MODULE=on GOOS=windows GOARCH=386 go1.12 build -ldflags -H=windowsgui
