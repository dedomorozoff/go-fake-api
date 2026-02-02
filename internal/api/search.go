package api

import (
	"net/http"
	"strings"

	"github.com/alexl/go-fake-api/internal/storage"
	"github.com/alexl/go-fake-api/internal/utils"
)

// Search выполняет поиск по миссиям и пилотам
func Search(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		if query == "" {
			utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
				"data": []interface{}{},
			})
			return
		}

		query = strings.ToLower(query)
		results := make([]map[string]interface{}, 0)

		// Поиск по миссиям
		missions := store.GetLunarMissions()
		for _, mission := range missions {
			// Проверяем название миссии
			if strings.Contains(strings.ToLower(mission.Name), query) {
				result := map[string]interface{}{
					"type":         "Миссия",
					"name":         mission.Name,
					"launch_date":  mission.LaunchDetails.LaunchDate,
					"landing_date": mission.LandingDetails.LandingDate,
					"crew":         mission.Spacecraft.Crew,
					"landing_site": mission.LandingDetails.LandingSite.Name,
				}
				results = append(results, result)
				continue
			}

			// Проверяем членов экипажа
			for _, crew := range mission.Spacecraft.Crew {
				if strings.Contains(strings.ToLower(crew.Name), query) {
					result := map[string]interface{}{
						"type":         "Миссия",
						"name":         mission.Name,
						"launch_date":  mission.LaunchDetails.LaunchDate,
						"landing_date": mission.LandingDetails.LandingDate,
						"crew":         mission.Spacecraft.Crew,
						"landing_site": mission.LandingDetails.LandingSite.Name,
					}
					results = append(results, result)
					break
				}
			}
		}

		response := map[string]interface{}{
			"data": results,
		}

		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}
