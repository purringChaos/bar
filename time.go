package main

import (
	"fmt"
	"time"

	"github.com/purringChaos/libKitteh/datetime"
)

type Time struct {
	Location *time.Location
	Prefix   string
}

type TimeWidget struct {
	s         *StatusBar
	times     []Time
	timeIndex int
}

func loadLocation(loc string) *time.Location {
	l, _ := time.LoadLocation(loc)
	return l
}

func NewTimeWidget(s *StatusBar, times []map[string]string) *TimeWidget {
	w := &TimeWidget{}
	w.s = s
	w.times = make([]Time, 0)

	for _, timeVal := range times {
		for key, value := range timeVal {
			w.times = append(w.times, Time{loadLocation(value), key})
		}
	}

	return w
}

func (w TimeWidget) InitialInfo() Info { return Info{"time", "none", "time", "#ffffff"} }

func (w TimeWidget) Name() string { return "time" }

func (w *TimeWidget) OnClick(e ClickEvent) {
	if w.timeIndex == len(w.times)-1 {
		w.timeIndex = 0
	} else {
		w.timeIndex = w.timeIndex + 1
	}
	w.Update()
}

func dateTimeToString(t time.Time) string {
	tdc := datetime.PrettyStruct(t, datetime.PrettyConfig{
		Use12HourTime:      true,
		RemoveEmptySeconds: false,
		HideSeconds:        false,
	})

	var timeString string
	timeString = timeString + Colour(RedColour, tdc.Hour)
	timeString = timeString + Colour(AccentLightColour, ":")
	timeString = timeString + Colour(OrangeColour, tdc.Minutes)
	if tdc.Seconds != "" {
		timeString = timeString + Colour(AccentLightColour, ":")
		timeString = timeString + Colour(YellowColour, tdc.Seconds)
	}
	timeString = timeString + Colour(AccentDarkColour, tdc.Ending)

	dateStr := fmt.Sprintf("%s %s %s%s %s %s %s %s %s %s",
		Colour(GreenColour, tdc.Weekday),
		Colour(PurpleColour, "the"),
		Colour(YellowColour, tdc.Day),
		Colour(AccentMediumColour, tdc.DayOrdinal),
		Colour(PurpleColour, "of"),
		Colour(RedColour, tdc.Month),
		Colour(PurpleColour, "in"),
		Colour(AccentLightColour, tdc.Year),
		Colour(PurpleColour, "at"),
		timeString,
	)

	return dateStr
}

func (w *TimeWidget) Update() {
	inLoc := w.times[w.timeIndex].Location
	if inLoc == nil {
		w.s.Add(Info{"time", "pango", w.times[w.timeIndex].Prefix + "Invalid Timezone", "#ffffff"})
		return
	}
	current := time.Now().In(inLoc)
	dateStr := dateTimeToString(current)
	w.s.Add(Info{"time", "pango", w.times[w.timeIndex].Prefix + dateStr, "#ffffff"})
}

func (w *TimeWidget) Start() {
	for {
		w.Update()
		time.Sleep(time.Millisecond * 500)
	}
}
