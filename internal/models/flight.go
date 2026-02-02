package models

// FlightCrewMember член экипажа полета
type FlightCrewMember struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// FlightLaunchSite место запуска полета
type FlightLaunchSite struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// FlightLaunchDetails детали запуска полета
type FlightLaunchDetails struct {
	LaunchDate string           `json:"launch_date"`
	LaunchSite FlightLaunchSite `json:"launch_site"`
}

// FlightLandingSite место посадки полета
type FlightLandingSite struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// FlightLandingDetails детали посадки полета
type FlightLandingDetails struct {
	LandingDate string            `json:"landing_date"`
	LandingSite FlightLandingSite `json:"landing_site"`
}

// Flight информация о полете
type Flight struct {
	Name           string               `json:"name"`
	CrewCapacity   int                  `json:"crew_capacity"`
	Cosmonaut      []FlightCrewMember   `json:"cosmonaut"`
	LaunchDetails  FlightLaunchDetails  `json:"launch_details"`
	LandingDetails FlightLandingDetails `json:"landing_details"`
}

// SpaceFlight космический рейс
type SpaceFlight struct {
	ID             int    `json:"id,omitempty"`
	FlightNumber   string `json:"flight_number"`
	Destination    string `json:"destination"`
	LaunchDate     string `json:"launch_date"`
	SeatsAvailable int    `json:"seats_available"`
}

// BookFlightRequest запрос на бронирование рейса
type BookFlightRequest struct {
	FlightNumber string `json:"flight_number"`
}

// FlightBooking бронирование рейса
type FlightBooking struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	FlightNumber string `json:"flight_number"`
}
