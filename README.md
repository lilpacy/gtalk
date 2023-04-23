## installation
```shell
go mod tidy
GOARCH=arm64 GOOS=darwin go build -o gtalk main.go # for m1 macos
sudo cp gtalk /usr/local/bin # adapt to the operating system
gtalk
```

## appendix
```shell
GOOS=linux GOARCH=amd64 go build -o gtalk main.go
GOOS=darwin GOARCH=amd64 go build -o gtalk main.go
GOOS=windows GOARCH=amd64 go build -o gtalk.exe main.go
```
