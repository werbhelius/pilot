package api

import (
	"encoding/json"
	"fmt"
	"github.com/werbhelius/pilot/model"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
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

const (
	CITY_API     = "https://raw.githubusercontent.com/werbhelius/Data/master/cityIdwithCoord.json"
	CURRENT_API  = "http://api.openweathermap.org/data/2.5/weather?id=%d&appid=57eeedbb8f9a8f7627af398802d741b0&units=metric&lang=%s"
	FORECAST_API = "http://api.openweathermap.org/data/2.5/forecast?id=%d&appid=57eeedbb8f9a8f7627af398802d741b0&units=metric&lang=%s"
)

func RequestCityId(cityName string) (int, error) {
	res, requestErr := http.Get(CITY_API)
	if requestErr != nil {
		return -1, fmt.Errorf("Unable to get city list")
	}
	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return -1, fmt.Errorf("Unable to read response body (%s): %v", CITY_API, bodyErr)
	}
	var resp []city
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return -1, fmt.Errorf("Unable to unmarshal response (%s): %v\nThe json body is: %s", CITY_API, jsonErr, string(body))
	}

	for _, city := range resp {
		if strings.ToUpper(city.Name) == strings.ToUpper(cityName) {
			return city.Id, nil
		}
	}

	return -1, fmt.Errorf("Unable to find city id by %s", cityName)

}

func CurrentWeather(cityid int, land string) model.Weather {
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

	var weather model.Weather
	weather.Coord = resp.Coord
	weather.Location = resp.Name
	weather.Now = model.Temperature{
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
	return weather
}
