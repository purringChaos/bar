package main

func main() {
	s := NewStatusBar()
	s.AddWidget(NewWeatherWidget(s, WeatherLocation, OpenWeatherAPIKey))
	s.AddWidget(NewBatteryWidget(s))
	s.AddWidget(NewTimeWidget(s, Times))
	s.Start()
}
