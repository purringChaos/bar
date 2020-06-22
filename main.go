package main

//import "time"

func main() {
	s := NewStatusBar()
	s.AddWidget(NewCPUWidget(s))
	//s.AddWidget(NewTextWidget(s, "OwO!", "#ffffff"))
	//s.AddWidget(NewRotatingTextWidget(s, TextRotations, time.Second * 5))
	s.AddWidget(NewWeatherWidget(s, WeatherLocation, OpenWeatherAPIKey))
	s.AddWidget(NewBatteryWidget(s))
	s.AddWidget(NewTimeWidget(s, Times))
	s.Start()
}
