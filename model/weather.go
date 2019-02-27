package model

import (
	"fmt"
	"time"
)

type Temperature struct {
	Time         time.Time
	WeatherCode  int
	WeatherDesc  string
	TempC        UnitTemp // °C
	TempC_max    UnitTemp // °C
	TempC_min    UnitTemp // °C
	Pressure     int      // hpa
	Humidity     int      //%
	Sunrise      time.Time
	Sunset       time.Time
	VisibilityM  float32 // m
	WindspeedMps float32 // m/s
	WindDegDesc  string
	RainOneHour  float32 // mm
	SnowOneHour  float32 // mm
	Cloudiness   string
}

type Day struct {
	Date     time.Time
	Weathers []Temperature
}

type Coord struct {
	Lat float32
	Lon float32
}

type Weather struct {
	Now      Temperature
	Forecast []Day
	Location string
	Coord    Coord
}

type UnitTemp float32

func (ut UnitTemp) FormatTemp() string {
	return fmt.Sprintf("%v°C", ut)
}
