package api

import (
	"encoding/json"
	"net/http"

	"github.com/alexl/go-fake-api/internal/models"
	"github.com/alexl/go-fake-api/internal/storage"
	"github.com/alexl/go-fake-api/internal/utils"
)

// CreateSpaceFlight создает новый космический рейс
func CreateSpaceFlight(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var flight models.SpaceFlight
		if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request", nil)
			return
		}

		// Валидация
		if validationErrors := utils.ValidateSpaceFlight(flight); len(validationErrors) > 0 {
			utils.RespondWithValidationError(w, validationErrors)
			return
		}

		// Создание рейса
		if err := store.CreateSpaceFlight(flight); err != nil {
			validationErrors := map[string][]string{
				"flight_number": {"flight with this number already exists"},
			}
			utils.RespondWithValidationError(w, validationErrors)
			return
		}

		response := map[string]interface{}{
			"data": map[string]interface{}{
				"code":    201,
				"message": "Космический полет создан",
			},
		}

		utils.RespondWithJSON(w, http.StatusCreated, response)
	}
}

// GetSpaceFlights возвращает список космических рейсов
func GetSpaceFlights(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flights := store.GetSpaceFlights()

		response := map[string]interface{}{
			"data": flights,
		}

		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}

// BookFlight бронирует рейс
func BookFlight(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		var req models.BookFlightRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request", nil)
			return
		}

		// Проверка наличия рейса
		flight, err := store.GetSpaceFlightByNumber(req.FlightNumber)
		if err != nil {
			code := 404
			utils.RespondWithError(w, http.StatusNotFound, "Not found", &code)
			return
		}

		// Проверка доступности мест
		if flight.SeatsAvailable <= 0 {
			validationErrors := map[string][]string{
				"flight_number": {"no seats available"},
			}
			utils.RespondWithValidationError(w, validationErrors)
			return
		}

		// Создание бронирования
		booking := models.FlightBooking{
			UserID:       user.ID,
			FlightNumber: req.FlightNumber,
		}

		if err := store.CreateFlightBooking(booking); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to book flight", nil)
			return
		}

		// Уменьшение количества доступных мест
		flight.SeatsAvailable--
		if err := store.UpdateSpaceFlight(*flight); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update flight", nil)
			return
		}

		response := map[string]interface{}{
			"data": map[string]interface{}{
				"code":    201,
				"message": "Рейс забронирован",
			},
		}

		utils.RespondWithJSON(w, http.StatusCreated, response)
	}
}
