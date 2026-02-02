# Space API - Краткое описание

## Что это?

Space API - это REST API на Go для работы с космическими данными. Проект демонстрирует современные практики разработки на Go и собирается в один исполняемый файл.

## Ключевые особенности

✅ **Один исполняемый файл** - все ресурсы встроены через `go:embed`  
✅ **Кросс-платформенность** - Windows, Linux, macOS, FreeBSD  
✅ **JWT аутентификация** - безопасная авторизация  
✅ **In-memory storage** - быстрая работа без БД  
✅ **Docker поддержка** - готовые Dockerfile и docker-compose  
✅ **GitHub Actions** - автоматическая сборка релизов  
✅ **Полная документация** - примеры для всех эндпоинтов  

## Технологии

- **Язык**: Go 1.21+
- **Роутинг**: Gorilla Mux
- **Аутентификация**: JWT (golang-jwt/jwt)
- **Хеширование**: bcrypt
- **Обработка изображений**: golang.org/x/image

## Быстрый старт

```bash
# Скачать и запустить
./space-api-linux-amd64

# Или собрать самостоятельно
go build -o space-api main.go
./space-api
```

Сервер запустится на http://localhost:8080

## API Endpoints

### Публичные
- `POST /registration` - Регистрация
- `POST /authorization` - Авторизация
- `GET /api/gagarin-flight` - Информация о полете Гагарина

### Защищенные (требуется Bearer токен)
- `GET /lunar-missions` - Список лунных миссий
- `POST /lunar-missions` - Добавить миссию
- `PATCH /lunar-missions/{id}` - Обновить миссию
- `DELETE /lunar-missions/{id}` - Удалить миссию
- `POST /space-flights` - Создать космический рейс
- `GET /space-flights` - Список рейсов
- `POST /book-flight` - Забронировать рейс
- `POST /lunar-watermark` - Создать водяной знак
- `GET /search?query=...` - Поиск

## Пример использования

```bash
# 1. Регистрация
curl -X POST http://localhost:8080/registration \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Alexey",
    "last_name": "Smirnov",
    "patronymic": "Ivanovich",
    "email": "user@prof.ru",
    "password": "paSSword1",
    "birth_date": "2001-02-15"
  }'

# 2. Авторизация
curl -X POST http://localhost:8080/authorization \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@prof.ru",
    "password": "paSSword1"
  }'

# 3. Получить миссии (с токеном)
curl http://localhost:8080/lunar-missions \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Сборка для разных платформ

```bash
# Использование Makefile
make release

# Или вручную
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o space-api.exe main.go
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o space-api-linux main.go
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o space-api-macos main.go
GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o space-api-freebsd main.go
```

## Структура проекта

```
go-fake-api/
├── main.go                 # Точка входа
├── internal/
│   ├── api/               # HTTP handlers
│   ├── middleware/        # Auth middleware
│   ├── models/            # Data models
│   ├── storage/           # In-memory storage
│   └── utils/             # Helpers
├── Makefile
├── Dockerfile
└── README.md
```

## Размер бинарника

- Без оптимизации: ~9 MB
- С оптимизацией (`-ldflags="-s -w"`): ~6.5 MB
- В Docker образе: ~15 MB (Alpine Linux)

## Производительность

- Запуск: < 1 сек
- Память: ~20-30 MB
- Concurrent запросы: поддерживаются

## Требования

**Для запуска готового бинарника**: Нет требований!  
**Для сборки**: Go 1.21+

## Лицензия

MIT License - свободное использование

## Документация

- [README.md](README.md) - Полная документация
- [API_TESTING.md](API_TESTING.md) - Примеры тестирования
- [CONTRIBUTING.md](CONTRIBUTING.md) - Руководство по вкладу

---

**Автор**: Space API Team  
**Год**: 2026  
**Статус**: Production Ready ✅
