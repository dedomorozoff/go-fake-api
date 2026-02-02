package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alexl/go-fake-api/internal/api"
	"github.com/alexl/go-fake-api/internal/middleware"
	"github.com/alexl/go-fake-api/internal/storage"
	"github.com/gorilla/mux"
)

func main() {
	// Инициализация хранилища
	store := storage.NewMemoryStorage()

	// Создание роутера
	r := mux.NewRouter()

	// Публичные эндпоинты
	r.HandleFunc("/registration", api.Registration(store)).Methods("POST")
	r.HandleFunc("/authorization", api.Authorization(store)).Methods("POST")
	r.HandleFunc("/api/gagarin-flight", api.GetGagarinFlight()).Methods("GET")

	// Защищенные эндпоинты
	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware(store))

	protected.HandleFunc("/logout", api.Logout(store)).Methods("GET")
	protected.HandleFunc("/flight", api.GetFlight()).Methods("GET")
	protected.HandleFunc("/lunar-missions", api.GetLunarMissions(store)).Methods("GET")
	protected.HandleFunc("/lunar-missions", api.AddLunarMission(store)).Methods("POST")
	protected.HandleFunc("/lunar-missions/{mission_id}", api.UpdateLunarMission(store)).Methods("PATCH")
	protected.HandleFunc("/lunar-missions/{mission_id}", api.DeleteLunarMission(store)).Methods("DELETE")
	protected.HandleFunc("/lunar-watermark", api.CreateLunarWatermark()).Methods("POST")
	protected.HandleFunc("/space-flights", api.CreateSpaceFlight(store)).Methods("POST")
	protected.HandleFunc("/space-flights", api.GetSpaceFlights(store)).Methods("GET")
	protected.HandleFunc("/book-flight", api.BookFlight(store)).Methods("POST")
	protected.HandleFunc("/search", api.Search(store)).Methods("GET")

	// Получение порта из переменной окружения или использование 8080 по умолчанию
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
