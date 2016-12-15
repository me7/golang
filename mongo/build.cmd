set GOOS=windows
set GOARCH=386
go build -ldflags "-s -w"
upx --best *.exe