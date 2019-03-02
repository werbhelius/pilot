package model

import (
	"fmt"
	"time"
)

type Temperature struct {
	Time         time.Time
	WeatherCode  UnitWeatherCode
	WeatherDesc  string
	TempC        UnitTemp // °C
	TempC_max    UnitTemp // °C
	TempC_min    UnitTemp // °C
	Pressure     float32  // hpa
	Humidity     float32  //%
	Sunrise      time.Time
	Sunset       time.Time
	VisibilityM  int     // m
	WindspeedMps float32 // m/s
	WindDegDesc  UnitWindDeg
	RainOneHour  float32 // mm
	SnowOneHour  float32 // mm
	Cloudiness   int
}

type Day struct {
	Date     time.Time
	Weathers []Temperature
}

type Coord struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type Location struct {
	Coord   Coord
	Id      int
	Name    string
	Country string
}

type Weather struct {
	Now      Temperature
	Forecast []Day
	Location Location
}

type UnitTemp float32

func (ut UnitTemp) FormatTemp() string {
	return fmt.Sprintf("%v°C", ut)
}

type UnitWindDeg float32

func (uw UnitWindDeg) FormatWindDeg() string {
	if uw >= 11.25 && uw < 33.75 {
		// NNE
		return "North-northeast"
	} else if uw >= 33.75 && uw < 56.25 {
		// NE
		return "North-east"
	} else if uw >= 56.25 && uw < 78.75 {
		// ENE
		return "East-northeast"
	} else if uw >= 78.75 && uw < 101.25 {
		// E
		return "East"
	} else if uw >= 101.25 && uw < 123.75 {
		// ESE
		return "East-southeast"
	} else if uw >= 123.75 && uw < 146.25 {
		// SE
		return "South-east"
	} else if uw >= 146.25 && uw < 168.75 {
		// SSE
		return "South-southeast"
	} else if uw >= 168.75 && uw < 191.25 {
		// S
		return "South"
	} else if uw >= 191.25 && uw < 213.75 {
		// SSW
		return "South-southwest"
	} else if uw >= 213.75 && uw < 236.25 {
		// SW
		return "South-west"
	} else if uw >= 236.25 && uw < 258.75 {
		// WSW
		return "West-southwest"
	} else if uw >= 258.75 && uw < 281.25 {
		// W
		return "West"
	} else if uw >= 281.25 && uw < 303.75 {
		// WNW
		return "West-northwest"
	} else if uw >= 303.75 && uw < 326.25 {
		// NW
		return "North-west"
	} else if uw >= 326.25 && uw < 348.75 {
		// NNW
		return "North-northwest"
	} else {
		// N
		return "North"
	}
}

type UnitWeatherCode int

func (uw UnitWeatherCode) FormatWeatherCode() string {
	if uw >= 200 && uw < 300 {
		return "⛈"
	} else if uw >= 300 && uw < 400 {
		return "🌦"
	} else if uw >= 500 && uw < 600 {
		return "🌧"
	} else if uw >= 600 && uw < 700 {
		return "❄️"
	} else if uw >= 700 && uw < 800 {
		return "🌫"
	} else if uw == 800 {
		return "☀️"
	} else if uw > 800 && uw < 900 {
		return "☁️"
	} else {
		return "🌝"
	}
}
