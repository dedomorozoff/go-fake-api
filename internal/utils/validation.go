package utils

import (
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/alexl/go-fake-api/internal/models"
)

// ValidateRegistration валидирует данные регистрации
func ValidateRegistration(req models.RegistrationRequest) map[string][]string {
	errors := make(map[string][]string)

	// Проверка first_name
	if req.FirstName == "" {
		errors["first_name"] = append(errors["first_name"], "field first_name can not be blank")
	} else if !isCapitalized(req.FirstName) {
		errors["first_name"] = append(errors["first_name"], "first letter must be uppercase")
	}

	// Проверка last_name
	if req.LastName == "" {
		errors["last_name"] = append(errors["last_name"], "field last_name can not be blank")
	} else if !isCapitalized(req.LastName) {
		errors["last_name"] = append(errors["last_name"], "first letter must be uppercase")
	}

	// Проверка patronymic
	if req.Patronymic == "" {
		errors["patronymic"] = append(errors["patronymic"], "field patronymic can not be blank")
	} else if !isCapitalized(req.Patronymic) {
		errors["patronymic"] = append(errors["patronymic"], "first letter must be uppercase")
	}

	// Проверка email
	if req.Email == "" {
		errors["email"] = append(errors["email"], "field email can not be blank")
	} else if !isValidEmail(req.Email) {
		errors["email"] = append(errors["email"], "invalid email format")
	}

	// Проверка password
	if req.Password == "" {
		errors["password"] = append(errors["password"], "field password can not be blank")
	} else if !isValidPassword(req.Password) {
		errors["password"] = append(errors["password"], "password must contain at least 3 characters, including at least one lowercase letter, one uppercase letter, and one digit")
	}

	// Проверка birth_date
	if req.BirthDate == "" {
		errors["birth_date"] = append(errors["birth_date"], "field birth_date can not be blank")
	} else if !isValidDate(req.BirthDate) {
		errors["birth_date"] = append(errors["birth_date"], "invalid date format, expected YYYY-MM-DD")
	}

	return errors
}

// ValidateAuthorization валидирует данные авторизации
func ValidateAuthorization(req models.AuthorizationRequest) map[string][]string {
	errors := make(map[string][]string)

	if req.Email == "" {
		errors["email"] = append(errors["email"], "field email can not be blank")
	}

	if req.Password == "" {
		errors["password"] = append(errors["password"], "field password can not be blank")
	}

	return errors
}

// ValidateLunarMission валидирует данные лунной миссии
func ValidateLunarMission(mission models.LunarMission) map[string][]string {
	errors := make(map[string][]string)

	if mission.Name == "" {
		errors["mission.name"] = append(errors["mission.name"], "field name can not be blank")
	}

	if mission.LaunchDetails.LaunchDate == "" {
		errors["mission.launch_details.launch_date"] = append(errors["mission.launch_details.launch_date"], "field launch_date can not be blank")
	} else if !isValidDate(mission.LaunchDetails.LaunchDate) {
		errors["mission.launch_details.launch_date"] = append(errors["mission.launch_details.launch_date"], "invalid date format")
	}

	if mission.LaunchDetails.LaunchSite.Name == "" {
		errors["mission.launch_details.launch_site.name"] = append(errors["mission.launch_details.launch_site.name"], "field name can not be blank")
	}

	if mission.LandingDetails.LandingDate == "" {
		errors["mission.landing_details.landing_date"] = append(errors["mission.landing_details.landing_date"], "field landing_date can not be blank")
	} else if !isValidDate(mission.LandingDetails.LandingDate) {
		errors["mission.landing_details.landing_date"] = append(errors["mission.landing_details.landing_date"], "invalid date format")
	}

	if mission.LandingDetails.LandingSite.Name == "" {
		errors["mission.landing_details.landing_site.name"] = append(errors["mission.landing_details.landing_site.name"], "field name can not be blank")
	}

	if mission.Spacecraft.CommandModule == "" {
		errors["mission.spacecraft.command_module"] = append(errors["mission.spacecraft.command_module"], "field command_module can not be blank")
	}

	if mission.Spacecraft.LunarModule == "" {
		errors["mission.spacecraft.lunar_module"] = append(errors["mission.spacecraft.lunar_module"], "field lunar_module can not be blank")
	}

	if len(mission.Spacecraft.Crew) == 0 {
		errors["mission.spacecraft.crew"] = append(errors["mission.spacecraft.crew"], "crew can not be empty")
	}

	return errors
}

// ValidateSpaceFlight валидирует данные космического рейса
func ValidateSpaceFlight(flight models.SpaceFlight) map[string][]string {
	errors := make(map[string][]string)

	if flight.FlightNumber == "" {
		errors["flight_number"] = append(errors["flight_number"], "field flight_number can not be blank")
	}

	if flight.Destination == "" {
		errors["destination"] = append(errors["destination"], "field destination can not be blank")
	}

	if flight.LaunchDate == "" {
		errors["launch_date"] = append(errors["launch_date"], "field launch_date can not be blank")
	} else if !isValidDate(flight.LaunchDate) {
		errors["launch_date"] = append(errors["launch_date"], "invalid date format")
	}

	if flight.SeatsAvailable <= 0 {
		errors["seats_available"] = append(errors["seats_available"], "seats_available must be greater than 0")
	}

	return errors
}

// ValidateWatermarkMessage валидирует сообщение для водяного знака
func ValidateWatermarkMessage(message string) map[string][]string {
	errors := make(map[string][]string)

	if message == "" {
		errors["message"] = append(errors["message"], "field message can not be blank")
	} else if len(message) < 10 || len(message) > 20 {
		errors["message"] = append(errors["message"], "message must be between 10 and 20 characters")
	}

	return errors
}

// isCapitalized проверяет, начинается ли строка с заглавной буквы
func isCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	firstRune := []rune(s)[0]
	return unicode.IsUpper(firstRune)
}

// isValidEmail проверяет валидность email
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// isValidPassword проверяет валидность пароля
func isValidPassword(password string) bool {
	if len(password) < 3 {
		return false
	}

	hasLower := false
	hasUpper := false
	hasDigit := false

	for _, char := range password {
		if unicode.IsLower(char) {
			hasLower = true
		}
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsDigit(char) {
			hasDigit = true
		}
	}

	return hasLower && hasUpper && hasDigit
}

// isValidDate проверяет валидность даты в формате YYYY-MM-DD
func isValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// NormalizeString нормализует строку (первая буква заглавная, остальные строчные)
func NormalizeString(s string) string {
	if len(s) == 0 {
		return s
	}
	
	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
