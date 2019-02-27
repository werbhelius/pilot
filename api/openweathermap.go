package api

import (
	"encoding/json"
	"fmt"
	"github.com/werbhelius/pilot/model"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type CityRequest struct {
	CityName string
}

type city struct {
	Id      int         `json:"id"`
	Name    string      `json:"name"`
	Country string      `json:"country"`
	Coord   model.Coord `json:"coord"`
}

type currentWeather struct {
	Coord   model.Coord `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Rain struct {
		OneHour float32 `json:"1h"`
	} `json:"rain"`
	Snow struct {
		OneHour float32 `json:"1h"`
	} `json:"snow"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Dt   int64  `json:"dt"`
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

type forecast struct {
	List []currentWeather `json:"list"`
}

const (
	CITY_API     = "https://raw.githubusercontent.com/werbhelius/Data/master/cityIdwithCoord.json"
	CURRENT_API  = "http://api.openweathermap.org/data/2.5/weather?id=%d&appid=57eeedbb8f9a8f7627af398802d741b0&units=metric&lang=%s"
	FORECAST_API = "http://api.openweathermap.org/data/2.5/forecast?id=%d&appid=57eeedbb8f9a8f7627af398802d741b0&units=metric&lang=%s"
)

func requestCityId(cityName string) model.Location {
	res, requestErr := http.Get(CITY_API)
	if requestErr != nil {
		log.Fatalf("Unable to get city list")
	}
	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Fatalf("Unable to read response body (%s): %v", CITY_API, bodyErr)
	}
	var resp []city
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		log.Fatalf("Unable to unmarshal response (%s): %v\nThe json body is: %s", CITY_API, jsonErr, string(body))
	}

	var location model.Location

	for _, city := range resp {
		if strings.ToUpper(city.Name) == strings.ToUpper(cityName) {
			location = model.Location{
				Id:    city.Id,
				Name:  city.Name,
				Coord: city.Coord,
			}
		}
	}

	return location
}

func nowWeather(cityid int, land string) model.Temperature {
	url := fmt.Sprintf(CURRENT_API, cityid, land)
	res, requestErr := http.Get(url)
	if requestErr != nil {
		log.Fatalln("Unable to get current weather")
	}
	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Fatalf("Unable to read response body (%s): %v", url, bodyErr)
	}
	var resp currentWeather
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		log.Fatalf("Unable to unmarshal response (%s): %v\nThe json body is: %s", url, jsonErr, string(body))
	}

	now := model.Temperature{
		Time:         time.Unix(resp.Dt, 0),
		WeatherCode:  resp.Weather[0].Id,
		WeatherDesc:  resp.Weather[0].Description,
		TempC:        model.UnitTemp(resp.Main.Temp),
		TempC_max:    model.UnitTemp(resp.Main.TempMax),
		TempC_min:    model.UnitTemp(resp.Main.TempMin),
		Pressure:     resp.Main.Pressure,
		Humidity:     resp.Main.Humidity,
		Sunrise:      time.Unix(resp.Sys.Sunrise, 0),
		Sunset:       time.Unix(resp.Sys.Sunset, 0),
		VisibilityM:  resp.Visibility,
		WindspeedMps: resp.Wind.Speed,
		WindDegDesc:  resp.Wind.Deg,
		RainOneHour:  resp.Rain.OneHour,
		SnowOneHour:  resp.Snow.OneHour,
		Cloudiness:   resp.Clouds.All,
	}
	return now
}

func forecastWeather(cityid int, land string) forecast {
	url := fmt.Sprintf(FORECAST_API, cityid, land)
	res, requestErr := http.Get(url)
	if requestErr != nil {
		log.Fatalln("Unable to get current weather")
	}
	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		log.Fatalf("Unable to read response body (%s): %v", url, bodyErr)
	}

	var forecast forecast
	jsonErr := json.Unmarshal(body, &forecast)
	if jsonErr != nil {
		log.Fatalf("Unable to unmarshal response (%s): %v\nThe json body is: %s", url, jsonErr, string(body))
	}

	return forecast
}

func Request(cityName string, land string) {
	location := requestCityId(cityName)
	if location.Id == 0 {
		log.Fatalf("Unable to find city id by %s", cityName)
	}

	var weather = model.Weather{}
	weather.Location = location

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		weather.Now = nowWeather(location.Id, land)
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		forecastWeather(location.Id, land)
		wg.Done()
	}(&wg)

	wg.Wait()

}
