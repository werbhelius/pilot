package main

import (
	"flag"
	"fmt"
	"github.com/werbhelius/pilot/api"
	"log"
)

func main() {

	// global flags
	city := flag.String(
		"city",
		"Beijing",
		"city must be request (you can find city name、id and coord in https://raw.githubusercontent.com/werbhelius/Data/master/cityIdwithCoord.json) ")

	lang := flag.String("lang", "zh_ch", "weather desc language ")

	flag.Parse()

	fmt.Println(*city)
	fmt.Println(*lang)

	// city requests
	id, err := api.RequestCityId(*city)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s's city id is %d\n", *city, id)

	// current weather
	weather := api.CurrentWeather(id, *lang)
	fmt.Printf("%s weather is %f\n", weather.Location, weather.Now.TempC)

}
