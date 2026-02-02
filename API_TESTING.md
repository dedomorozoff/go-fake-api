# API Testing Examples

Этот файл содержит примеры тестирования всех эндпоинтов API.

## 1. Регистрация пользователя

```bash
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
```

Ожидаемый ответ (201):
```json
{
  "data": {
    "user": {
      "name": "Smirnov Alexey Ivanovich",
      "email": "user@prof.ru"
    },
    "code": 201,
    "message": "Пользователь создан"
  }
}
```

## 2. Авторизация

```bash
curl -X POST http://localhost:8080/authorization \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@prof.ru",
    "password": "paSSword1"
  }'
```

Ожидаемый ответ (200):
```json
{
  "data": {
    "user": {
      "id": 1,
      "name": "Smirnov Alexey Ivanovich",
      "birth_date": "2001-02-15",
      "email": "user@prof.ru"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Сохраните токен для дальнейших запросов!**

```bash
export TOKEN="your-token-here"
```

## 3. Получение информации о полете Гагарина (публичный)

```bash
curl http://localhost:8080/api/gagarin-flight
```

## 4. Получение информации о полете Аполлон-11 (защищенный)

```bash
curl http://localhost:8080/flight \
  -H "Authorization: Bearer $TOKEN"
```

## 5. Получение списка лунных миссий

```bash
curl http://localhost:8080/lunar-missions \
  -H "Authorization: Bearer $TOKEN"
```

## 6. Добавление новой лунной миссии

```bash
curl -X POST http://localhost:8080/lunar-missions \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "mission": {
      "name": "Аполлон-12",
      "launch_details": {
        "launch_date": "1969-11-14",
        "launch_site": {
          "name": "Космический центр имени Кеннеди",
          "location": {
            "latitude": "28.5721",
            "longitude": "-80.648"
          }
        }
      },
      "landing_details": {
        "landing_date": "1969-11-19",
        "landing_site": {
          "name": "Океан Бурь",
          "coordinates": {
            "latitude": "-3.0",
            "longitude": "-23.4"
          }
        }
      },
      "spacecraft": {
        "command_module": "Янки Клиппер",
        "lunar_module": "Интрепид",
        "crew": [
          {
            "name": "Чарльз Конрад",
            "role": "Командир"
          },
          {
            "name": "Алан Бин",
            "role": "Пилот лунного модуля"
          },
          {
            "name": "Ричард Гордон",
            "role": "Пилот командного модуля"
          }
        ]
      }
    }
  }'
```

## 7. Обновление лунной миссии

```bash
curl -X PATCH http://localhost:8080/lunar-missions/1 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "mission": {
      "name": "Аполлон-11 Обновленный",
      "launch_details": {
        "launch_date": "1969-07-16",
        "launch_site": {
          "name": "Космический центр имени Кеннеди",
          "location": {
            "latitude": "28.5721",
            "longitude": "-80.648"
          }
        }
      },
      "landing_details": {
        "landing_date": "1969-07-20",
        "landing_site": {
          "name": "Море спокойствия",
          "coordinates": {
            "latitude": "0.674",
            "longitude": "23.472"
          }
        }
      },
      "spacecraft": {
        "command_module": "Колумбия",
        "lunar_module": "Орел",
        "crew": [
          {
            "name": "Нил Армстронг",
            "role": "Командир"
          },
          {
            "name": "Базз Олдрин",
            "role": "Пилот лунного модуля"
          },
          {
            "name": "Майкл Коллинз",
            "role": "Пилот командного модуля"
          }
        ]
      }
    }
  }'
```

## 8. Удаление лунной миссии

```bash
curl -X DELETE http://localhost:8080/lunar-missions/3 \
  -H "Authorization: Bearer $TOKEN"
```

## 9. Создание космического рейса

```bash
curl -X POST http://localhost:8080/space-flights \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "flight_number": "СФ-201",
    "destination": "Луна",
    "launch_date": "2026-03-15",
    "seats_available": 10
  }'
```

## 10. Получение списка космических рейсов

```bash
curl http://localhost:8080/space-flights \
  -H "Authorization: Bearer $TOKEN"
```

## 11. Бронирование рейса

```bash
curl -X POST http://localhost:8080/book-flight \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "flight_number": "СФ-103"
  }'
```

## 12. Поиск по миссиям

```bash
# Поиск по названию миссии
curl "http://localhost:8080/search?query=Аполлон" \
  -H "Authorization: Bearer $TOKEN"

# Поиск по имени космонавта
curl "http://localhost:8080/search?query=Армстронг" \
  -H "Authorization: Bearer $TOKEN"
```

## 13. Создание водяного знака на изображении

```bash
# Создайте тестовое изображение или используйте существующее
curl -X POST http://localhost:8080/lunar-watermark \
  -H "Authorization: Bearer $TOKEN" \
  -F "fileimage=@/path/to/your/image.jpg" \
  -F "message=Профессионалы 2024" \
  --output watermarked.jpg
```

## 14. Выход из системы

```bash
curl http://localhost:8080/logout \
  -H "Authorization: Bearer $TOKEN"
```

## Тестирование ошибок

### Попытка доступа без токена (403)

```bash
curl http://localhost:8080/lunar-missions
```

Ожидаемый ответ:
```json
{
  "message": "Login failed"
}
```

### Попытка получить несуществующий ресурс (404)

```bash
curl -X DELETE http://localhost:8080/lunar-missions/999 \
  -H "Authorization: Bearer $TOKEN"
```

Ожидаемый ответ:
```json
{
  "message": "Not found",
  "code": 404
}
```

### Ошибка валидации (422)

```bash
curl -X POST http://localhost:8080/registration \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "",
    "last_name": "",
    "patronymic": "",
    "email": "invalid-email",
    "password": "weak",
    "birth_date": "invalid-date"
  }'
```

Ожидаемый ответ:
```json
{
  "error": {
    "code": 422,
    "message": "Validation error",
    "errors": {
      "first_name": ["field first_name can not be blank"],
      "last_name": ["field last_name can not be blank"],
      "patronymic": ["field patronymic can not be blank"],
      "email": ["invalid email format"],
      "password": ["password must contain at least 3 characters..."],
      "birth_date": ["invalid date format, expected YYYY-MM-DD"]
    }
  }
}
```
