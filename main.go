package main

import "time"

func main() {
	s := NewStatusBar()
	s.AddWidget(NewTextWidget(s, "<b>"+Colour("red", "☭")+"</b>", "#ff0000"))
	s.AddWidget(NewRotatingTextWidget(s, []string{BgColour("black", Colour("red", "Anar")) + BgColour("red", Colour("black", "chy!")), BgColour("red", Colour("black", "Anar")) + BgColour("black", Colour("red", "chy!"))}, time.Minute/2))
	s.AddWidget(NewMemoryWidget(s))
	s.AddWidget(NewCPUWidget(s))
	s.AddWidget(NewWeatherWidget(s, WeatherLocation, OpenWeatherAPIKey))
	s.AddWidget(NewBatteryWidget(s))
	s.AddWidget(NewTimeWidget(s, Times))
	s.Start()
}
