package models

// Location представляет географические координаты
type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Coordinates представляет координаты
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// LaunchSite представляет место запуска
type LaunchSite struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

// LandingSite представляет место посадки
type LandingSite struct {
	Name        string      `json:"name"`
	Coordinates Coordinates `json:"coordinates"`
}

// LaunchDetails содержит детали запуска
type LaunchDetails struct {
	LaunchDate string     `json:"launch_date"`
	LaunchSite LaunchSite `json:"launch_site"`
}

// LandingDetails содержит детали посадки
type LandingDetails struct {
	LandingDate string      `json:"landing_date"`
	LandingSite LandingSite `json:"landing_site"`
}

// CrewMember представляет члена экипажа
type CrewMember struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// Spacecraft представляет космический корабль
type Spacecraft struct {
	CommandModule string       `json:"command_module"`
	LunarModule   string       `json:"lunar_module"`
	Crew          []CrewMember `json:"crew"`
}

// LunarMission представляет лунную миссию
type LunarMission struct {
	ID             int            `json:"id,omitempty"`
	Name           string         `json:"name"`
	LaunchDetails  LaunchDetails  `json:"launch_details"`
	LandingDetails LandingDetails `json:"landing_details"`
	Spacecraft     Spacecraft     `json:"spacecraft"`
}

// LunarMissionWrapper обертка для миссии
type LunarMissionWrapper struct {
	Mission LunarMission `json:"mission"`
}

// GagarinFlightDuration длительность полета
type GagarinFlightDuration struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

// GagarinSpacecraft космический корабль Гагарина
type GagarinSpacecraft struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	CrewCapacity int    `json:"crew_capacity"`
}

// GagarinLaunchSite место запуска Гагарина
type GagarinLaunchSite struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

// GagarinLaunchDetails детали запуска Гагарина
type GagarinLaunchDetails struct {
	LaunchDate string            `json:"launch_date"`
	LaunchSite GagarinLaunchSite `json:"launch_site"`
}

// GagarinLandingCoordinates координаты посадки
type GagarinLandingCoordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// GagarinLandingSite место посадки
type GagarinLandingSite struct {
	Name        string                    `json:"name"`
	Country     string                    `json:"country"`
	Coordinates GagarinLandingCoordinates `json:"coordinates"`
}

// GagarinLandingDetails детали посадки
type GagarinLandingDetails struct {
	ParachuteLanding  bool    `json:"parachute_landing"`
	ImpactVelocityMps int     `json:"impact_velocity_mps"`
}

// GagarinLanding информация о посадке
type GagarinLanding struct {
	Date    string                `json:"date"`
	Site    GagarinLandingSite    `json:"site"`
	Details GagarinLandingDetails `json:"details"`
}

// GagarinBio биография Гагарина
type GagarinBio struct {
	EarlyLife  string `json:"early_life"`
	Career     string `json:"career"`
	PostFlight string `json:"post_flight"`
}

// GagarinCosmonaut информация о космонавте
type GagarinCosmonaut struct {
	Name      string     `json:"name"`
	Birthdate string     `json:"birthdate"`
	Rank      string     `json:"rank"`
	Bio       GagarinBio `json:"bio"`
}

// GagarinMission миссия Гагарина
type GagarinMission struct {
	Name           string                `json:"name"`
	LaunchDetails  GagarinLaunchDetails  `json:"launch_details"`
	FlightDuration GagarinFlightDuration `json:"flight_duration"`
	Spacecraft     GagarinSpacecraft     `json:"spacecraft"`
}

// GagarinFlight полная информация о полете Гагарина
type GagarinFlight struct {
	Mission    GagarinMission   `json:"mission"`
	Landing    GagarinLanding   `json:"landing"`
	Cosmonaut  GagarinCosmonaut `json:"cosmonaut"`
}
