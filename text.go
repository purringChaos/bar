package main


type TextWidget struct {
	s *StatusBar
	text string
        colour string
        sep bool
        sbw int
}

func NewTextWidget(s *StatusBar, text, colour string) TextWidget {
	w := TextWidget{}
	w.s = s
	w.text = text
	w.colour = colour
	return w
}

func (w TextWidget) InitialInfo() Info {
	return Info{w.text, "pango", w.text, w.colour}
}

func (w TextWidget) Name() string {
	return w.text
}

func (w TextWidget) OnClick(e ClickEvent) {
	return
}

func (w TextWidget) Start() {
    return
}