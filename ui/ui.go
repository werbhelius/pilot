package ui

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/werbhelius/pilot/api"
	"github.com/werbhelius/pilot/model"
	"sync"
	"time"
)

func Render(cityName string, land string) {
	var done = new(bool)
	*done = false

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		printLoading(done)
		wg.Done()
	}(&wg)

	weather := new(model.Weather)

	go func(wg *sync.WaitGroup) {
		*weather = api.Request(cityName, land)
		*done = true
		wg.Done()
	}(&wg)

	wg.Wait()
	printCityInfo(weather.Location)
	printNow(weather.Now)

}

func printLoading(done *bool) {
	var index = new(int)
	yellow := color.New(color.FgYellow).SprintFunc()
	for {
		*index++
		if *index == 1 {
			fmt.Printf("\r%s", yellow("Waiting."))
		}
		if *index == 2 {
			fmt.Printf("\r%s", yellow("Waiting.."))
		}
		if *index == 3 {
			fmt.Printf("\r%s", yellow("Waiting..."))
			*index = 0
		}
		time.Sleep(1 * time.Second / 2)
		if *done {
			fmt.Printf("\r%s", "")
			break
		}
	}
}

func printCityInfo(location model.Location) {
	green := color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	fmt.Printf("Weather in %s,%s (%f,%f)\n\n", green(location.Name), green(location.Country), location.Coord.Lat, location.Coord.Lon)
}

func printNow(now model.Temperature) {
	blue := color.New(color.FgHiBlue)
	green := color.New(color.FgGreen).Add(color.Bold).SprintFunc()

	_, _ = blue.Print("Current Weather: ")
	fmt.Printf("%s %s\n", green(now.WeatherDesc), green(now.TempC.FormatTemp()))

	_, _ = blue.Print("Wind: ")
	fmt.Printf("%s %sm/s\n", green(now.WindDegDesc.FormatWindDeg()), green(now.WindspeedMps))

	_, _ = blue.Print("Cloudiness: ")

	_, _ = blue.Print("Pressure: ")
	fmt.Printf("%shpa\n", green(now.Pressure))

	_, _ = blue.Print("Humidity: ")
	fmt.Printf("%s %%\n", green(now.Humidity))

	_, _ = blue.Print("Sunrise: ")
	fmt.Printf("%s\n", green(now.Sunrise.Format("15:04")))

	_, _ = blue.Print("Sunset: ")
	fmt.Printf("%s\n", green(now.Sunset.Format("15:04")))

}
