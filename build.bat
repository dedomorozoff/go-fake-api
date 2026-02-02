@echo off
REM Скрипт для сборки релизов под все платформы (Windows)

echo Building releases for all platforms...

REM Создание директории для релизов
if not exist releases mkdir releases

REM Windows
echo Building for Windows (amd64)...
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o releases/space-api-windows-amd64.exe main.go

REM Linux
echo Building for Linux (amd64)...
set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o releases/space-api-linux-amd64 main.go

echo Building for Linux (arm64)...
set GOOS=linux
set GOARCH=arm64
go build -ldflags="-s -w" -o releases/space-api-linux-arm64 main.go

REM macOS
echo Building for macOS (amd64)...
set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-s -w" -o releases/space-api-macos-amd64 main.go

echo Building for macOS (arm64/M1)...
set GOOS=darwin
set GOARCH=arm64
go build -ldflags="-s -w" -o releases/space-api-macos-arm64 main.go

REM FreeBSD
echo Building for FreeBSD (amd64)...
set GOOS=freebsd
set GOARCH=amd64
go build -ldflags="-s -w" -o releases/space-api-freebsd-amd64 main.go

echo Building for FreeBSD (arm64)...
set GOOS=freebsd
set GOARCH=arm64
go build -ldflags="-s -w" -o releases/space-api-freebsd-arm64 main.go

echo Done! Binaries are in the 'releases' directory.
dir releases
