package storage

import (
	"sync"

	"github.com/alexl/go-fake-api/internal/models"
)

// Storage интерфейс хранилища
type Storage interface {
	// Users
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
	UpdateUserToken(userID int, token string) error

	// Lunar Missions
	GetLunarMissions() []models.LunarMission
	GetLunarMissionByID(id int) (*models.LunarMission, error)
	AddLunarMission(mission models.LunarMission) error
	UpdateLunarMission(id int, mission models.LunarMission) error
	DeleteLunarMission(id int) error

	// Space Flights
	CreateSpaceFlight(flight models.SpaceFlight) error
	GetSpaceFlights() []models.SpaceFlight
	GetSpaceFlightByNumber(flightNumber string) (*models.SpaceFlight, error)
	UpdateSpaceFlight(flight models.SpaceFlight) error

	// Flight Bookings
	CreateFlightBooking(booking models.FlightBooking) error
	GetUserBookings(userID int) []models.FlightBooking
}

// MemoryStorage хранилище в памяти
type MemoryStorage struct {
	users          map[int]*models.User
	usersByEmail   map[string]*models.User
	usersByToken   map[string]*models.User
	lunarMissions  map[int]models.LunarMission
	spaceFlights   map[string]models.SpaceFlight
	flightBookings []models.FlightBooking
	userIDCounter  int
	missionIDCounter int
	bookingIDCounter int
	mu             sync.RWMutex
}

// NewMemoryStorage создает новое хранилище в памяти
func NewMemoryStorage() *MemoryStorage {
	storage := &MemoryStorage{
		users:          make(map[int]*models.User),
		usersByEmail:   make(map[string]*models.User),
		usersByToken:   make(map[string]*models.User),
		lunarMissions:  make(map[int]models.LunarMission),
		spaceFlights:   make(map[string]models.SpaceFlight),
		flightBookings: make([]models.FlightBooking, 0),
		userIDCounter:  1,
		missionIDCounter: 1,
		bookingIDCounter: 1,
	}
	
	// Инициализация начальных данных
	storage.initDefaultMissions()
	storage.initDefaultFlights()
	
	return storage
}

// initDefaultMissions инициализирует миссии по умолчанию
func (s *MemoryStorage) initDefaultMissions() {
	missions := []models.LunarMission{
		{
			ID:   1,
			Name: "Аполлон-11",
			LaunchDetails: models.LaunchDetails{
				LaunchDate: "1969-07-16",
				LaunchSite: models.LaunchSite{
					Name: "Космический центр имени Кеннеди",
					Location: models.Location{
						Latitude:  "28.5721000",
						Longitude: "-80.6480000",
					},
				},
			},
			LandingDetails: models.LandingDetails{
				LandingDate: "1969-07-20",
				LandingSite: models.LandingSite{
					Name: "Море спокойствия",
					Coordinates: models.Coordinates{
						Latitude:  "0.6740000",
						Longitude: "23.4720000",
					},
				},
			},
			Spacecraft: models.Spacecraft{
				CommandModule: "Колумбия",
				LunarModule:   "Орел",
				Crew: []models.CrewMember{
					{Name: "Нил Армстронг", Role: "Командир"},
					{Name: "Базз Олдрин", Role: "Пилот лунного модуля"},
					{Name: "Майкл Коллинз", Role: "Пилот командного модуля"},
				},
			},
		},
		{
			ID:   2,
			Name: "Аполлон-17",
			LaunchDetails: models.LaunchDetails{
				LaunchDate: "1972-12-07",
				LaunchSite: models.LaunchSite{
					Name: "Космический центр имени Кеннеди",
					Location: models.Location{
						Latitude:  "28.5721000",
						Longitude: "-80.6480000",
					},
				},
			},
			LandingDetails: models.LandingDetails{
				LandingDate: "1972-12-19",
				LandingSite: models.LandingSite{
					Name: "Телец-Литтров",
					Coordinates: models.Coordinates{
						Latitude:  "20.1908000",
						Longitude: "30.7717000",
					},
				},
			},
			Spacecraft: models.Spacecraft{
				CommandModule: "Америка",
				LunarModule:   "Челленджер",
				Crew: []models.CrewMember{
					{Name: "Евгений Сернан", Role: "Командир"},
					{Name: "Харрисон Шмитт", Role: "Пилот лунного модуля"},
					{Name: "Рональд Эванс", Role: "Пилот командного модуля"},
				},
			},
		},
	}

	for _, mission := range missions {
		s.lunarMissions[mission.ID] = mission
	}
	s.missionIDCounter = 3
}

// initDefaultFlights инициализирует рейсы по умолчанию
func (s *MemoryStorage) initDefaultFlights() {
	flights := []models.SpaceFlight{
		{
			FlightNumber:   "СФ-103",
			Destination:    "Марс",
			LaunchDate:     "2025-05-15",
			SeatsAvailable: 2,
		},
		{
			FlightNumber:   "СФ-105",
			Destination:    "Юпитер",
			LaunchDate:     "2024-06-01",
			SeatsAvailable: 3,
		},
	}

	for _, flight := range flights {
		s.spaceFlights[flight.FlightNumber] = flight
	}
}
