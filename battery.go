package main

import (
	"fmt"
	"path/filepath"
	"time"
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
	return Info{"battery", "pango", "battery", TextColour}
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
		batInfo := Colour(AccentLightColour, Bold("bat")) + " "
		capacity := readFileAsFloat(capacityPath)
		status := readFileAsString(statusPath)

		colour := TextColour
		if status == "Charging" {
			descriptor = Colour(GreenColour, "(C)")
			posNegIndicator = "+"
			colour = GreenColour
		} else if status == "Discharging" {
			if capacity > 50 {
				descriptor = Colour(OrangeColour, "(D)")
			} else {
				descriptor = Colour(RedColour, "(D)")
			}
			posNegIndicator = "-"
		} else if status == "Unknown" {
			descriptor = Colour(YellowColour, "(D)")
			posNegIndicator = "?"
		}

		if status != "Charging" {
			if capacity > 80 {
				colour = GreenColour
			} else if capacity > 60 {
				colour = YellowColour
			} else if capacity > 40 {
				colour = OrangeColour
			} else {
				colour = RedColour
			}
		}

		batInfo = batInfo + descriptor
		batInfo = batInfo + " " + Colour(colour, fmt.Sprintf("%d", int(capacity)) + Colour(AccentDarkColour, "%"))

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
			batInfo = batInfo + " " + Colour(PurpleColour, fmt.Sprintf("%s%.2fW", posNegIndicator, watts))
		}

		w.s.Add(Info{"battery", "pango", batInfo, TextColour})
		time.Sleep(time.Millisecond * 400)
	}
}
