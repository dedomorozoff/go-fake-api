# Используем многоэтапную сборку для минимизации размера образа
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go mod файлы
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o space-api main.go

# Финальный образ
FROM alpine:latest

# Устанавливаем CA сертификаты для HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/space-api .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./space-api"]
