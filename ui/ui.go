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
	printTodayHourWeather(weather.Forecast[0])
	printForecastWeather(weather.Forecast)

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
	fmt.Printf("🌝---------------------Weather in %s,%s (%f,%f)---------------------🌝\n\n", green(location.Name), green(location.Country), location.Coord.Lat, location.Coord.Lon)
}

//noinspection GoUnhandledErrorResult
func printNow(now model.Temperature) {
	blue := color.New(color.FgHiBlue)
	green := color.New(color.FgGreen).Add(color.Bold)

	blue.Print("Current Weather: ")
	green.Printf("%s %s ", now.WeatherDesc, now.TempC.FormatTemp())
	fmt.Print(now.WeatherCode.FormatWeatherCode())
	fmt.Print("\n")

	blue.Print("Wind: ")
	green.Printf("%s %gm/s\n", now.WindDegDesc.FormatWindDeg(), now.WindspeedMps)

	blue.Print("Cloudiness: ")
	green.Printf("%d%%\n", now.Cloudiness)

	blue.Print("Pressure: ")
	green.Printf("%ghpa\n", now.Pressure)

	blue.Print("Humidity: ")
	green.Printf("%g%%\n", now.Humidity)

	blue.Print("Sunrise: ")
	green.Printf("%s\n", now.Sunrise.Format("15:04"))

	blue.Print("Sunset: ")
	green.Printf("%s\n\n", now.Sunset.Format("15:04"))

}

//noinspection GoUnhandledErrorResult
func printTodayHourWeather(day model.Day) {
	blue := color.New(color.FgHiBlue)
	green := color.New(color.FgGreen).Add(color.Bold)

	blue.Print("Today Hourly weather: \n")
	for _, weather := range day.Weathers {
		blue.Printf("%s: ", weather.Time.Format("15:04"))
		green.Printf("%s %s ", weather.WeatherDesc, weather.TempC.FormatTemp())
		fmt.Print(weather.WeatherCode.FormatWeatherCode())
		fmt.Print("\n")
	}
	blue.Print("\n")
}

//noinspection GoUnhandledErrorResult
func printForecastWeather(forecast []model.Day) {
	blue := color.New(color.FgHiBlue)
	green := color.New(color.FgGreen).Add(color.Bold)

	blue.Print("Forecast weather: \n")
	for index, day := range forecast {
		if index == 0 {
			continue
		}
		blue.Printf("%s %s weather: \n", day.Date.Format("2006-01-02"), day.Date.Weekday().String())
		for _, weather := range day.Weathers {
			blue.Printf("%s: ", weather.Time.Format("15:04"))
			green.Printf("%s %s %s\n", weather.WeatherDesc, weather.TempC.FormatTemp(), weather.WeatherCode.FormatWeatherCode())
		}
		blue.Print("\n")
	}
}
