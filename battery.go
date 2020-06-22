package main

import (
	"fmt"
	"time"
	"path/filepath"
)

type BatteryWidget struct {
	s *StatusBar
}

func NewBatteryWidget(s *StatusBar) BatteryWidget {
	w := BatteryWidget{}
	w.s = s
	return w
}

func (w BatteryWidget) InitialInfo() Info {
	return Info{"battery", "none", "battery", "#ffffff", }
}

func (w BatteryWidget) Name() string {
	return "battery"
}

func (w BatteryWidget) OnClick(e ClickEvent) {
	return
}


func globGetFirst(fp string) string {
	paths, err := filepath.Glob(fp)
	if err != nil {
		return ""
	} else if len(paths) == 0 {
		return ""
	} else {
		return paths[0]
	}
}

func (w BatteryWidget) Start() {
	statusPath := globGetFirst("/sys/class/power_supply/*/status")
	powerNowPath := globGetFirst("/sys/class/power_supply/*/power_now")
	capacityPath := globGetFirst("/sys/class/power_supply/*/capacity")
	currentNowPath := globGetFirst("/sys/class/power_supply/*/current_now")
	voltageNowPath := globGetFirst("/sys/class/power_supply/*/voltage_now")

	var watts float64
	var descriptor string
	var posNegIndicator string


	for {
		canGetWatts := false
		batInfo := "bat "
		capacity := readFileAsFloat(capacityPath)
		status := readFileAsString(statusPath)

		colour := "#ffffff"
		if status == "Charging" {
			descriptor = "C"
			posNegIndicator = "+"
			colour = "#00ff00"
		} else if status == "Discharging" {
			descriptor = "D"
			posNegIndicator = "-"
		}  else if status == "Unknown" {
			descriptor = "U"
			posNegIndicator = "?"
		}

		if status != "Charging" {
			if capacity > 70 {
				colour = "#00ff00"
			} else if capacity > 40 {
				colour = "#FFA500"
			} else {
				colour = "#ff0000"
			}
		}

		batInfo = batInfo + fmt.Sprintf("(%s)", descriptor)
		batInfo = batInfo + fmt.Sprintf(" %d%%", int(capacity))

		if powerNowPath != "" {
			watts = readFileAsFloat(powerNowPath) / 1000000
			canGetWatts = true
		} else if currentNowPath != "" && voltageNowPath != "" {
			currentNow := readFileAsFloat(currentNowPath) / 1000000
			voltageNow := readFileAsFloat(voltageNowPath) / 1000000
			if currentNow == 0 || voltageNow == 0 {
				canGetWatts = false
			} else {
				watts = (currentNow * voltageNow)
				canGetWatts = true
			}
		}

		if canGetWatts {
			batInfo = batInfo + fmt.Sprintf(" %s%.2fW", posNegIndicator, watts)
		}

		w.s.Add(Info{"battery", "none", batInfo, colour, })
		time.Sleep(time.Millisecond * 400)
	}
}
