package main

import (
	"encoding/hex"
	"math/rand"
	"time"
	"sync"
)

func genRandString() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

type RotatingTextWidget struct {
	s        *StatusBar
	texts    []string
	index    int
	randID   string
	duration time.Duration
	mutex sync.Mutex
}

func NewRotatingTextWidget(s *StatusBar, texts []string, duration time.Duration) *RotatingTextWidget {
	w := &RotatingTextWidget{}
	w.s = s
	w.texts = texts
	w.randID = genRandString()
	w.duration = duration
	return w
}

func (w RotatingTextWidget) InitialInfo() Info {
	return Info{w.randID, "pango", "", "#ffffff"}
}

func (w RotatingTextWidget) Name() string {
	return w.randID
}

func (w *RotatingTextWidget) OnClick(e ClickEvent) {
	w.rotate()
}

func (w *RotatingTextWidget) rotate() {
	w.mutex.Lock()
	if w.index == len(w.texts)-1 {
		w.index = 0
	} else {
		w.index = w.index + 1
	}
	w.mutex.Unlock()
	w.update()
}

func (w *RotatingTextWidget) update() {
	w.mutex.Lock()
	w.s.Add(Info{w.randID, "pango", w.texts[w.index], "#ffffff"})
	w.mutex.Unlock()
}

func (w *RotatingTextWidget) Start() {
	for {
		w.update()
		w.rotate()
		time.Sleep(w.duration)
	}
}
