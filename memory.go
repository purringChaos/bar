package main

import (
	"fmt"
	psMem "github.com/shirou/gopsutil/mem"
	"time"
	"github.com/dustin/go-humanize"
)

type MemoryWidget struct {
	s        *StatusBar
	statuses []string
	index    int
}

func formatMemoryPercent(percent float64) (out string) {
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
	return Colour(percentColour, fmt.Sprintf("%3.0f", percent)) + Colour(AccentDarkColour, "%")
}


func NewMemoryWidget(s *StatusBar) *MemoryWidget {
	w := &MemoryWidget{}
	w.s = s
	return w
}

func (w MemoryWidget) InitialInfo() Info {
	return Info{"Memory", "pango", "Memory", "#ffffff"}
}

func (w MemoryWidget) Name() string {
	return "Memory"
}

func (w *MemoryWidget) OnClick(e ClickEvent) {
	if w.index == len(w.statuses)-1 {
		w.index = 0
	} else {
		w.index = w.index + 1
	}
	w.update()
}

/*func (w *MemoryWidget) updateAllMemory() {
	for {
		percent, _ := psMem.Percent(time.Second/4, false)
		w.statuses[0] = formatMemoryPercent(0, percent[0])
	}
	w.update()
}

func (w *MemoryWidget) updatePerMemory() {
	for {
		percents, _ := psMem.Percent(time.Second, true)
		for i, percent := range percents {
			w.statuses[i+1] = formatMemoryPercent(i+1, percent)
		}
		w.update()
	}
}*/

func (w MemoryWidget) updateIndex(i int) {
	if w.index == i {
		w.update()
	}
}

func (w MemoryWidget) update() {
	w.s.Add(Info{"Memory", "pango", w.statuses[w.index], "#ffffff"})
}

func (w *MemoryWidget) vmupdater() {
	for {
		vm, _ := psMem.VirtualMemory()
		swp, _ := psMem.SwapMemory()
		w.statuses[0] = Colour(AccentLightColour, Bold("mem ")) + formatMemoryPercent(vm.UsedPercent)
		w.updateIndex(0)
		w.statuses[1] = Colour(AccentLightColour, Bold("swap ")) + formatMemoryPercent(swp.UsedPercent)
		w.updateIndex(1)
		w.statuses[2] = Colour(AccentLightColour, Bold("mem free ")) + humanize.Bytes(vm.Free)
		w.updateIndex(2)
		w.statuses[3] = Colour(AccentLightColour, Bold("swap free ")) + humanize.Bytes(swp.Free)
		w.updateIndex(3)
		w.statuses[4] = Colour(AccentLightColour, Bold("mem used ")) + humanize.Bytes(vm.Used)
		w.updateIndex(4)
		w.statuses[5] = Colour(AccentLightColour, Bold("swap used ")) + humanize.Bytes(swp.Used)
		w.updateIndex(5)
		time.Sleep(time.Second / 4)
	}
}


func (w *MemoryWidget) Start() {
	w.statuses = make([]string, 6)
	go w.vmupdater()	
}
