package main

import (
	"github.com/purringChaos/libKitteh/datetime"
	"time"
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

func (w TimeWidget) InitialInfo() Info {
	return Info{"time", "none", "time", "#ffffff"}
}

func (w TimeWidget) Name() string {
	return "time"
}

func (w *TimeWidget) OnClick(e ClickEvent) {
	if w.timeIndex == len(w.times)-1 {
		w.timeIndex = 0
	} else {
		w.timeIndex = w.timeIndex + 1
	}
	w.Update()
	return
}

func (w *TimeWidget) Update() {
	inLoc := w.times[w.timeIndex].Location
	if inLoc == nil {
		w.s.Add(Info{"time", "none", w.times[w.timeIndex].Prefix + "Invalid Timezone", "#ffffff"})
		return
	}
	current := time.Now().In(inLoc)
	dateStr := datetime.Pretty(current, datetime.PrettyConfig{
		Use12HourTime:      true,
		RemoveEmptySeconds: false,
		HideSeconds:        false,
	})
	w.s.Add(Info{"time", "none", w.times[w.timeIndex].Prefix + dateStr, "#ffffff"})
}

func (w *TimeWidget) Start() {
	for {
		w.Update()
		time.Sleep(time.Millisecond * 500)
	}
}
