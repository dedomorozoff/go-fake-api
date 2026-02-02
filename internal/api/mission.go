package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alexl/go-fake-api/internal/models"
	"github.com/alexl/go-fake-api/internal/storage"
	"github.com/alexl/go-fake-api/internal/utils"
	"github.com/gorilla/mux"
)

// GetLunarMissions возвращает список лунных миссий
func GetLunarMissions(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		missions := store.GetLunarMissions()

		// Формируем ответ в требуемом формате
		response := make([]map[string]interface{}, 0)
		for _, mission := range missions {
			response = append(response, map[string]interface{}{
				"mission": mission,
			})
		}

		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}

// AddLunarMission добавляет новую лунную миссию
func AddLunarMission(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var wrapper models.LunarMissionWrapper
		if err := json.NewDecoder(r.Body).Decode(&wrapper); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request", nil)
			return
		}

		// Валидация
		if validationErrors := utils.ValidateLunarMission(wrapper.Mission); len(validationErrors) > 0 {
			utils.RespondWithValidationError(w, validationErrors)
			return
		}

		// Добавление миссии
		if err := store.AddLunarMission(wrapper.Mission); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to add mission", nil)
			return
		}

		response := map[string]interface{}{
			"data": map[string]interface{}{
				"code":    201,
				"message": "Миссия добавлена",
			},
		}

		utils.RespondWithJSON(w, http.StatusCreated, response)
	}
}

// UpdateLunarMission обновляет лунную миссию
func UpdateLunarMission(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		missionID, err := strconv.Atoi(vars["mission_id"])
		if err != nil {
			code := 404
			utils.RespondWithError(w, http.StatusNotFound, "Not found", &code)
			return
		}

		// Проверка существования миссии
		if _, err := store.GetLunarMissionByID(missionID); err != nil {
			code := 404
			utils.RespondWithError(w, http.StatusNotFound, "Not found", &code)
			return
		}

		var wrapper models.LunarMissionWrapper
		if err := json.NewDecoder(r.Body).Decode(&wrapper); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request", nil)
			return
		}

		// Валидация
		if validationErrors := utils.ValidateLunarMission(wrapper.Mission); len(validationErrors) > 0 {
			utils.RespondWithValidationError(w, validationErrors)
			return
		}

		// Обновление миссии
		if err := store.UpdateLunarMission(missionID, wrapper.Mission); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update mission", nil)
			return
		}

		response := map[string]interface{}{
			"data": map[string]interface{}{
				"code":    200,
				"message": "Миссия обновлена",
			},
		}

		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}

// DeleteLunarMission удаляет лунную миссию
func DeleteLunarMission(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		missionID, err := strconv.Atoi(vars["mission_id"])
		if err != nil {
			code := 404
			utils.RespondWithError(w, http.StatusNotFound, "Not found", &code)
			return
		}

		// Проверка существования миссии
		if _, err := store.GetLunarMissionByID(missionID); err != nil {
			code := 404
			utils.RespondWithError(w, http.StatusNotFound, "Not found", &code)
			return
		}

		// Удаление миссии
		if err := store.DeleteLunarMission(missionID); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete mission", nil)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
