package main

func main() {
	s := NewStatusBar()
        s.AddWidget(NewTextWidget(s, "<span foreground=\"#F7A8B8\">Tr</span><span foreground=\"#55CDFC\">an</span><span foreground=\"white\">s Ri</span><span foreground=\"#55CDFC\">gh</span><span foreground=\"#F7A8B8\">ts</span>!", "#ffffff"))
	s.AddWidget(NewWeatherWidget(s, WeatherLocation, OpenWeatherAPIKey))
	s.AddWidget(NewBatteryWidget(s))
	s.AddWidget(NewTimeWidget(s, Times))
	s.Start()
}
