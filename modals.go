package main

// CityWeather holds data required for the webapp
type CityWeather struct {
	City     string
	Temp     float64
	Pressure float64
	Humidity int
	TempMaxC float64
	TempMinC float64
	Desc     string
}
