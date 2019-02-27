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
	Pressure     float32  // hpa
	Humidity     float32  //%
	Sunrise      time.Time
	Sunset       time.Time
	VisibilityM  int     // m
	WindspeedMps float32 // m/s
	WindDegDesc  float32
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
	Coord Coord
	Id    int
	Name  string
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

type UnitWindDeg string

func (uw UnitWindDeg) FormatWindDeg() {

}
