package ui

import (
	"fmt"
	"github.com/werbhelius/pilot/api"
)

func Render(cityName string, land string) {

	weather := api.Request(cityName, land)

	fmt.Printf("Weather in %s,%s (%f,%f)\n", weather.Location.Name, weather.Location.Country, weather.Location.Coord.Lat, weather.Location.Coord.Lon)

}

func format() {

}
