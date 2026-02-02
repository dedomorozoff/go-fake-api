package models

import (
	"time"
)

// User представляет пользователя системы
type User struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Patronymic string    `json:"patronymic"`
	Email      string    `json:"email"`
	Password   string    `json:"-"` // Не отдаем пароль в JSON
	BirthDate  string    `json:"birth_date"`
	Token      string    `json:"-"`
	CreatedAt  time.Time `json:"-"`
}

// GetFullName возвращает полное имя пользователя
func (u *User) GetFullName() string {
	return u.LastName + " " + u.FirstName + " " + u.Patronymic
}

// RegistrationRequest структура запроса регистрации
type RegistrationRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	BirthDate  string `json:"birth_date"`
}

// AuthorizationRequest структура запроса авторизации
type AuthorizationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
