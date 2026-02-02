#!/bin/bash

# Скрипт для сборки релизов под все платформы

echo "Building releases for all platforms..."

# Создание директории для релизов
mkdir -p releases

# Windows
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o releases/space-api-windows-amd64.exe main.go

# Linux
echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o releases/space-api-linux-amd64 main.go

echo "Building for Linux (arm64)..."
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o releases/space-api-linux-arm64 main.go

# macOS
echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o releases/space-api-macos-amd64 main.go

echo "Building for macOS (arm64/M1)..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o releases/space-api-macos-arm64 main.go

# FreeBSD
echo "Building for FreeBSD (amd64)..."
GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o releases/space-api-freebsd-amd64 main.go

echo "Building for FreeBSD (arm64)..."
GOOS=freebsd GOARCH=arm64 go build -ldflags="-s -w" -o releases/space-api-freebsd-arm64 main.go

echo "Done! Binaries are in the 'releases' directory."
ls -lh releases/
