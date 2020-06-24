package main

import (
	"fmt"
	"time"

	psCpu "github.com/shirou/gopsutil/cpu"
)

type CPUWidget struct {
	s        *StatusBar
	statuses []string
	index    int
}

func formatCPUPercent(cpuNum int, percent float64) (out string) {
	// If cpuNum == 0 then it is the adverage of all cores
	// as the rest of the cores are 1 indexed

	var percentColour string
	if percent > 80 {
		percentColour = RedColour
	} else if percent > 60 {
		percentColour = OrangeColour
	} else if percent > 30 {
		percentColour = YellowColour
	} else {
		percentColour = GreenColour
	}

	if cpuNum == 0 {
		out = out + Colour(AccentLightColour, Bold("cpu"))
		out = out + " "
		out = out + Colour(percentColour, fmt.Sprintf("%3.0f", percent))
		out = out + Colour(AccentDarkColour, "%")
	} else {
		out = out + Colour(AccentLightColour, Bold("cpu")+Colour(YellowColour, "#")+fmt.Sprintf("%d", cpuNum))
		out = out + " "
		out = out + Colour(percentColour, fmt.Sprintf("%3.0f", percent))
		out = out + Colour(AccentDarkColour, "%")
	}
	return out
}

func NewCPUWidget(s *StatusBar) *CPUWidget {
	w := &CPUWidget{}
	w.s = s
	return w
}

func (w CPUWidget) InitialInfo() Info {
	return Info{"cpu", "pango", "cpu", "#ffffff"}
}

func (w CPUWidget) Name() string {
	return "cpu"
}

func (w *CPUWidget) OnClick(e ClickEvent) {
	if w.index == len(w.statuses)-1 {
		w.index = 0
	} else {
		w.index = w.index + 1
	}
	w.update()
}

func (w *CPUWidget) updateAllCPU() {
	for {
		percent, _ := psCpu.Percent(time.Second/4, false)
		w.statuses[0] = formatCPUPercent(0, percent[0])
		w.update()
	}
}

func (w *CPUWidget) updatePerCPU() {
	for {
		percents, _ := psCpu.Percent(time.Second, true)
		for i, percent := range percents {
			w.statuses[i+1] = formatCPUPercent(i+1, percent)
			w.update()
		}
	}
}

func (w CPUWidget) update() {
	w.s.Add(Info{"cpu", "pango", w.statuses[w.index], "#ffffff"})
}

func (w *CPUWidget) Start() {
	cpuCount, _ := psCpu.Counts(true)
	w.statuses = make([]string, cpuCount+1)
	go w.updateAllCPU()
	go w.updatePerCPU()
}
