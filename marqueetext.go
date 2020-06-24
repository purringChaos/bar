package main

import "time"

type MarqueeTextWidget struct {
	s     *StatusBar
	text  string
	width int
	id    string
}

func NewMarqueeTextWidget(s *StatusBar, id string, text string, width int) *MarqueeTextWidget {
	w := &MarqueeTextWidget{}
	w.s = s
	w.text = text
	w.id = id
	w.width = width
	return w
}

func (w *MarqueeTextWidget) InitialInfo() Info {
	return Info{w.id, "pango", "", TextColour}
}

func (w *MarqueeTextWidget) Name() string {
	return w.id
}

func (w MarqueeTextWidget) OnClick(e ClickEvent) {
	return
}

func (w MarqueeTextWidget) Start() {
	for (len(w.text) % w.width) != 1 {
		w.text = w.text + " "
	}
	w.text = w.text + w.text[0:w.width-1]
	i := 0
	for {
		w.s.Add(Info{w.id, "pango", w.text[i : i+w.width], TextColour})
		if i+w.width == len(w.text) {
			i = 0
		} else {
			i = i + 1
		}
		time.Sleep(time.Millisecond * 120)
	}
}
