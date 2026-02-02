.PHONY: build run test clean release help

# Переменные
BINARY_NAME=space-api
MAIN_FILE=main.go
BUILD_DIR=releases

# Цвета для вывода
GREEN=\033[0;32m
NC=\033[0m # No Color

help: ## Показать справку
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  ${GREEN}%-15s${NC} %s\n", $$1, $$2}'

build: ## Собрать проект для текущей платформы
	@echo "Building ${BINARY_NAME}..."
	@go build -ldflags="-s -w" -o ${BINARY_NAME} ${MAIN_FILE}
	@echo "Build complete: ${BINARY_NAME}"

run: ## Запустить сервер
	@echo "Starting server..."
	@go run ${MAIN_FILE}

test: ## Запустить тесты
	@echo "Running tests..."
	@go test -v ./...

clean: ## Очистить собранные файлы
	@echo "Cleaning..."
	@rm -f ${BINARY_NAME} ${BINARY_NAME}.exe
	@rm -rf ${BUILD_DIR}
	@echo "Clean complete"

deps: ## Установить зависимости
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies installed"

release: ## Собрать релизы для всех платформ
	@echo "Building releases for all platforms..."
	@mkdir -p ${BUILD_DIR}
	
	@echo "Building for Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${BINARY_NAME}-windows-amd64.exe ${MAIN_FILE}
	
	@echo "Building for Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${BINARY_NAME}-linux-amd64 ${MAIN_FILE}
	
	@echo "Building for Linux (arm64)..."
	@GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${BINARY_NAME}-linux-arm64 ${MAIN_FILE}
	
	@echo "Building for macOS (amd64)..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${BINARY_NAME}-macos-amd64 ${MAIN_FILE}
	
	@echo "Building for macOS (arm64/M1)..."
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${BINARY_NAME}-macos-arm64 ${MAIN_FILE}
	
	@echo "Building for FreeBSD (amd64)..."
	@GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${BINARY_NAME}-freebsd-amd64 ${MAIN_FILE}
	
	@echo "Building for FreeBSD (arm64)..."
	@GOOS=freebsd GOARCH=arm64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${BINARY_NAME}-freebsd-arm64 ${MAIN_FILE}
	
	@echo "Release build complete!"
	@ls -lh ${BUILD_DIR}/

dev: ## Запустить в режиме разработки с hot reload (требует air)
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not installed. Install with: go install github.com/cosmtrek/air@latest"; \
		echo "Running without hot reload..."; \
		go run ${MAIN_FILE}; \
	fi

docker-build: ## Собрать Docker образ
	@echo "Building Docker image..."
	@docker build -t ${BINARY_NAME}:latest .

docker-run: ## Запустить в Docker
	@echo "Running in Docker..."
	@docker run -p 8080:8080 ${BINARY_NAME}:latest
