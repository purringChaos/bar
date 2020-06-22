package main

import (
	"fmt"
	psCpu "github.com/shirou/gopsutil/cpu"
	"time"
)

type CPUWidget struct {
	s        *StatusBar
	statuses []string
	index    int
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
		w.statuses[0] = fmt.Sprintf("cpu %3.0f%%", percent[0])
	}
	w.update()
}

func (w *CPUWidget) updatePerCPU() {
	for {
		percents, _ := psCpu.Percent(time.Second, true)
		for i, percent := range percents {
			w.statuses[i+1] = fmt.Sprintf("cpu#%d %3.0f%%", i+1, percent)
		}
		w.update()
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
