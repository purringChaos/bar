# Bar

This is just a simple i3bar/swaybar status_command thing that I use.

To use it simply create a secrets.go file with contents like:
```go
package main

const OpenWeatherAPIKey = "your api key"
const WeatherLocation = "your location"

var Times = map[string]string{
	"": "Europe/London",
}
```
and then compile it and set your status_command in conf.
Feel free to disable any module in main.go and fork it for your own use.

You may also wish to change the acceptable temp ranges for colour in weather.go.