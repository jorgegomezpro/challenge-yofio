rmdir /S /Q bin
set GOOS=linux
go build -ldflags="-s -w" -o bin/main main.go

cd bin

%GOPATH%/bin/win-go-zipper.exe -o main.zip ./