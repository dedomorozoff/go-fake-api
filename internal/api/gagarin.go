package api

import (
	"net/http"

	"github.com/alexl/go-fake-api/internal/models"
	"github.com/alexl/go-fake-api/internal/utils"
)

// GetGagarinFlight возвращает информацию о полете Гагарина
func GetGagarinFlight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gagarinFlight := models.GagarinFlight{
			Mission: models.GagarinMission{
				Name: "Восток 1",
				LaunchDetails: models.GagarinLaunchDetails{
					LaunchDate: "1961-04-12",
					LaunchSite: models.GagarinLaunchSite{
						Name: "Космодром Байконур",
						Location: models.Location{
							Latitude:  "45.9650000",
							Longitude: "63.3050000",
						},
					},
				},
				FlightDuration: models.GagarinFlightDuration{
					Hours:   1,
					Minutes: 48,
				},
				Spacecraft: models.GagarinSpacecraft{
					Name:         "Восток 3KA",
					Manufacturer: "OKB-1",
					CrewCapacity: 1,
				},
			},
			Landing: models.GagarinLanding{
				Date: "1961-04-12",
				Site: models.GagarinLandingSite{
					Name:    "Смеловка",
					Country: "СССР",
					Coordinates: models.GagarinLandingCoordinates{
						Latitude:  "51.2700000",
						Longitude: "45.9970000",
					},
				},
				Details: models.GagarinLandingDetails{
					ParachuteLanding:  true,
					ImpactVelocityMps: 7,
				},
			},
			Cosmonaut: models.GagarinCosmonaut{
				Name:      "Юрий Гагарин",
				Birthdate: "1934-03-09",
				Rank:      "Старший лейтенант",
				Bio: models.GagarinBio{
					EarlyLife:  "Родился в Клушино, Россия.",
					Career:     "Отобран в отряд космонавтов в 1960 году...",
					PostFlight: "Стал международным героем.",
				},
			},
		}

		response := map[string]interface{}{
			"data": []models.GagarinFlight{gagarinFlight},
		}

		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}

// GetFlight возвращает информацию о полете Аполлон-11
func GetFlight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flight := models.Flight{
			Name:         "Аполлон-11",
			CrewCapacity: 3,
			Cosmonaut: []models.FlightCrewMember{
				{Name: "Нил Армстронг", Role: "Командир"},
				{Name: "Базз Олдрин", Role: "Пилот лунного модуля"},
				{Name: "Майкл Коллинз", Role: "Пилот командного модуля"},
			},
			LaunchDetails: models.FlightLaunchDetails{
				LaunchDate: "1969-07-16",
				LaunchSite: models.FlightLaunchSite{
					Name:      "Космический центр имени Кеннеди",
					Latitude:  "28.5721000",
					Longitude: "-80.6480000",
				},
			},
			LandingDetails: models.FlightLandingDetails{
				LandingDate: "1969-07-20",
				LandingSite: models.FlightLandingSite{
					Name:      "Море спокойствия",
					Latitude:  "0.6740000",
					Longitude: "23.4720000",
				},
			},
		}

		response := map[string]interface{}{
			"data": flight,
		}

		utils.RespondWithJSON(w, http.StatusOK, response)
	}
}
