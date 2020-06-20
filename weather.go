package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Weather struct {
	Main string `json:main`
}
type MainWeather struct {
	Temp float32 `json:"temp"`
}

type WeatherData struct {
	Weathers []Weather   `json:"weather"`
	Main     MainWeather `json:"main"`
}

type WeatherWidget struct {
	s *StatusBar
	location string
	apikey string
}

func NewWeatherWidget(s *StatusBar, loc, apikey string) WeatherWidget {
	w := WeatherWidget{}
	w.s = s
	w.location = loc
	w.apikey = apikey
	return w
}

func (w WeatherWidget) InitialInfo() Info {
	return Info{"weather", "none", "weather", "#ffffff"}
}

func (w WeatherWidget) Name() string {
	return "weather"
}

func (w WeatherWidget) OnClick(e ClickEvent) {
	return
}

func (w WeatherWidget) Start() {
	for {
		res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q="+w.location+"&appid="+w.apikey+"&units=metric")
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}
		var weather WeatherData
		err = json.NewDecoder(res.Body).Decode(&weather)
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}

		i := Info{"weather", "none", fmt.Sprintf("%d°C %s", int(weather.Main.Temp), weather.Weathers[0].Main), "#ffffff"}
		// My preferred temp ranges.
		if weather.Main.Temp >= 20 {
			i.Colour = "#ff0000"
		} else if weather.Main.Temp <= 18 {
			i.Colour = "#00ff00"
		}

		w.s.Add(i)
		time.Sleep(time.Minute * 15)
	}
}
