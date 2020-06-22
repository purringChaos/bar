package main

import "time"

func main() {
	s := NewStatusBar()
	s.AddWidget(NewTextWidget(s, "OwO!", "#ffffff"))
	s.AddWidget(NewRotatingTextWidget(s, TextRotations, time.Millisecond * 1))
	s.AddWidget(NewWeatherWidget(s, WeatherLocation, OpenWeatherAPIKey))
	s.AddWidget(NewBatteryWidget(s))
	s.AddWidget(NewTimeWidget(s, Times))
	s.Start()
}
