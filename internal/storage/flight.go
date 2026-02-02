package storage

import (
	"errors"

	"github.com/alexl/go-fake-api/internal/models"
)

// CreateSpaceFlight создает новый космический рейс
func (s *MemoryStorage) CreateSpaceFlight(flight models.SpaceFlight) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.spaceFlights[flight.FlightNumber]; exists {
		return errors.New("flight with this number already exists")
	}

	s.spaceFlights[flight.FlightNumber] = flight
	return nil
}

// GetSpaceFlights возвращает все космические рейсы
func (s *MemoryStorage) GetSpaceFlights() []models.SpaceFlight {
	s.mu.RLock()
	defer s.mu.RUnlock()

	flights := make([]models.SpaceFlight, 0, len(s.spaceFlights))
	for _, flight := range s.spaceFlights {
		flights = append(flights, flight)
	}

	return flights
}

// GetSpaceFlightByNumber получает рейс по номеру
func (s *MemoryStorage) GetSpaceFlightByNumber(flightNumber string) (*models.SpaceFlight, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	flight, exists := s.spaceFlights[flightNumber]
	if !exists {
		return nil, errors.New("flight not found")
	}

	return &flight, nil
}

// UpdateSpaceFlight обновляет рейс
func (s *MemoryStorage) UpdateSpaceFlight(flight models.SpaceFlight) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.spaceFlights[flight.FlightNumber]; !exists {
		return errors.New("flight not found")
	}

	s.spaceFlights[flight.FlightNumber] = flight
	return nil
}

// CreateFlightBooking создает бронирование рейса
func (s *MemoryStorage) CreateFlightBooking(booking models.FlightBooking) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	booking.ID = s.bookingIDCounter
	s.bookingIDCounter++
	s.flightBookings = append(s.flightBookings, booking)

	return nil
}

// GetUserBookings получает бронирования пользователя
func (s *MemoryStorage) GetUserBookings(userID int) []models.FlightBooking {
	s.mu.RLock()
	defer s.mu.RUnlock()

	bookings := make([]models.FlightBooking, 0)
	for _, booking := range s.flightBookings {
		if booking.UserID == userID {
			bookings = append(bookings, booking)
		}
	}

	return bookings
}
