package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const weatherURL = "http://api.openweathermap.org/data/2.5/weather"
const appid string = "5771e2497ba720d90ccb2563f50a3b4f"

// Coord ...
type Coord struct {
	Lon float64
	Lat float64
}

// Weather ..
type Weather struct {
	ID          int
	Main        string
	Description string
}

// Main ...
type Main struct {
	Temp     float64
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressure int
	Humidity int
}

// Wind ...
type Wind struct {
	Speed int
	Deg   int
}

// Clouds ...
type Clouds struct {
	All int
}

// Sys ...
type Sys struct {
	Type    int
	ID      int
	Country string
	Sunrise int64
	Sunset  int64
}

// AllInfo ...
type AllInfo struct {
	Coord      *Coord
	Weather    []*Weather
	Base       string
	Main       *Main
	visibility string
	Wind       *Wind
	Clouds     *Clouds
	Dt         int64
	Sys        *Sys
	Name       string
}

// GetWeatherInfo ...
func GetWeatherInfo(city string, units string) (*AllInfo, error) {
	resp, err := http.Get(weatherURL + "?q=" + city + "&units=" + units + "&" + "appid=" + appid)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса : %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка запроса : %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка записи : %s", err)
	}
	var coord AllInfo
	if err := json.Unmarshal(body, &coord); err != nil {
		return nil, fmt.Errorf("ошибка декодирования %s", err)
	}
	return &coord, nil
}
