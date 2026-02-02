# ⚡ Быстрый старт - Space API

## 🚀 За 5 минут до запуска!

### Вариант 1: Готовый бинарник (самый быстрый)

```bash
# 1. Скачайте для вашей платформы
# Windows
curl -LO https://github.com/your-repo/releases/latest/download/space-api-windows-amd64.exe

# Linux
curl -LO https://github.com/your-repo/releases/latest/download/space-api-linux-amd64
chmod +x space-api-linux-amd64

# macOS
curl -LO https://github.com/your-repo/releases/latest/download/space-api-macos-amd64
chmod +x space-api-macos-amd64

# 2. Запустите
./space-api-linux-amd64  # или .exe для Windows

# 3. Готово! API работает на http://localhost:8080
```

### Вариант 2: Docker (рекомендуется)

```bash
# 1. Клонируйте репозиторий
git clone https://github.com/your-repo/go-fake-api.git
cd go-fake-api

# 2. Запустите
docker-compose up -d

# 3. Готово! API работает на http://localhost:8080
```

### Вариант 3: Сборка из исходников

```bash
# 1. Клонируйте
git clone https://github.com/your-repo/go-fake-api.git
cd go-fake-api

# 2. Соберите
go build -o space-api main.go

# 3. Запустите
./space-api

# 4. Готово! API работает на http://localhost:8080
```

---

## 🧪 Первый тест

```bash
# Проверьте, что API работает
curl http://localhost:8080/api/gagarin-flight

# Должен вернуть JSON с информацией о полете Гагарина
```

---

## 📝 Первые запросы

### 1. Регистрация пользователя

```bash
curl -X POST http://localhost:8080/registration \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Alexey",
    "last_name": "Smirnov",
    "patronymic": "Ivanovich",
    "email": "test@example.com",
    "password": "Test123",
    "birth_date": "2000-01-01"
  }'
```

**Ответ:**
```json
{
  "data": {
    "user": {
      "name": "Smirnov Alexey Ivanovich",
      "email": "test@example.com"
    },
    "code": 201,
    "message": "Пользователь создан"
  }
}
```

### 2. Авторизация

```bash
curl -X POST http://localhost:8080/authorization \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "Test123"
  }'
```

**Ответ:**
```json
{
  "data": {
    "user": {
      "id": 1,
      "name": "Smirnov Alexey Ivanovich",
      "birth_date": "2000-01-01",
      "email": "test@example.com"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**💡 Сохраните токен!** Он понадобится для следующих запросов.

### 3. Получение лунных миссий

```bash
# Замените YOUR_TOKEN на токен из предыдущего шага
curl http://localhost:8080/lunar-missions \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Ответ:**
```json
[
  {
    "mission": {
      "name": "Аполлон-11",
      "launch_details": { ... },
      "landing_details": { ... },
      "spacecraft": { ... }
    }
  },
  ...
]
```

---

## 🎯 Что дальше?

### Изучите документацию

1. **[README.md](README.md)** - Полная документация
2. **[API_TESTING.md](API_TESTING.md)** - Все примеры API
3. **[DEPLOYMENT.md](DEPLOYMENT.md)** - Развертывание в продакшн

### Попробуйте другие эндпоинты

```bash
# Создать космический рейс
curl -X POST http://localhost:8080/space-flights \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "flight_number": "SF-001",
    "destination": "Луна",
    "launch_date": "2026-12-31",
    "seats_available": 10
  }'

# Забронировать рейс
curl -X POST http://localhost:8080/book-flight \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"flight_number": "SF-001"}'

# Поиск по миссиям
curl "http://localhost:8080/search?query=Аполлон" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## 🛠️ Полезные команды

### Разработка

```bash
# Запуск с hot reload (требует air)
make dev

# Обычный запуск
make run

# Запуск тестов
make test
```

### Сборка

```bash
# Сборка для текущей платформы
make build

# Сборка для всех платформ
make release

# Очистка
make clean
```

### Docker

```bash
# Запуск
docker-compose up -d

# Остановка
docker-compose down

# Логи
docker-compose logs -f

# Перезапуск
docker-compose restart
```

---

## 📚 Все эндпоинты

### Публичные (без токена)

| Метод | URL | Описание |
|-------|-----|----------|
| POST | `/registration` | Регистрация |
| POST | `/authorization` | Авторизация |
| GET | `/api/gagarin-flight` | Информация о Гагарине |

### Защищенные (требуется Bearer токен)

| Метод | URL | Описание |
|-------|-----|----------|
| GET | `/logout` | Выход |
| GET | `/flight` | Информация о полете |
| GET | `/lunar-missions` | Список миссий |
| POST | `/lunar-missions` | Добавить миссию |
| PATCH | `/lunar-missions/{id}` | Обновить миссию |
| DELETE | `/lunar-missions/{id}` | Удалить миссию |
| POST | `/lunar-watermark` | Водяной знак |
| POST | `/space-flights` | Создать рейс |
| GET | `/space-flights` | Список рейсов |
| POST | `/book-flight` | Забронировать |
| GET | `/search?query=...` | Поиск |

---

## ❓ Частые вопросы

### Как изменить порт?

```bash
export PORT=3000  # Linux/macOS
set PORT=3000     # Windows
./space-api
```

### Где хранятся данные?

В памяти приложения. При перезапуске данные теряются.

### Как остановить сервер?

```bash
# Ctrl+C в терминале

# Или найти и убить процесс
ps aux | grep space-api
kill <PID>
```

### Нужна ли база данных?

Нет! API использует in-memory хранилище.

---

## 🐛 Проблемы?

### Порт занят

```bash
# Измените порт
export PORT=8081
./space-api
```

### Не работает curl

Используйте Postman, Insomnia или браузер для GET запросов.

### Ошибка "Login failed"

Проверьте, что:
1. Вы зарегистрировались
2. Токен правильный
3. Используете `Authorization: Bearer TOKEN`

---

## 📞 Помощь

- **Документация:** [README.md](README.md)
- **Примеры:** [API_TESTING.md](API_TESTING.md)
- **Issues:** GitHub Issues
- **Email:** your-email@example.com

---

## ✅ Чек-лист

- [ ] API запущен
- [ ] Зарегистрировал пользователя
- [ ] Получил токен
- [ ] Протестировал эндпоинты
- [ ] Прочитал документацию
- [ ] Готов к разработке!

---

**🎉 Поздравляем! Вы готовы работать с Space API!**

Следующий шаг: [README.md](README.md) для полной документации
