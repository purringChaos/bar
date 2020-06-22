package main

import (
	"encoding/json"
	"fmt"
	"github.com/purringChaos/libKitteh/filesystem"
	"os"
	"reflect"
	"sync"
	"regexp"
)

func debugLog(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

var colourRegex = regexp.MustCompile("(#[0-9a-fA-F]{6}){(.*?)}")
var italicRegex = regexp.MustCompile(`\$(i|italic|italics){(.*)}`)
var boldRegex = regexp.MustCompile(`\$(b|bold){(.*)}`)

func rewriteInfoText(s string) string {
	s = colourRegex.ReplaceAllString(s, "<span foreground=\"$1\">$2</span>")
	s = italicRegex.ReplaceAllString(s, "<i>$2</i>")
	s = boldRegex.ReplaceAllString(s, "<b>$2</b>")
	return s
}

type Info struct {
	Name   string `json:"name"`
	Markup string `json:"markup"`
	Text   string `json:"full_text"`
	Colour string `json:"color,omitempty"`
}

type ClickEvent struct {
	Name string `json:"name"`
}

func waitForever() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

type AllInfo []Info

func readFileAsFloat(fn string) float64 {
	data, _ := filesystem.ReadFloat(fn)
	return data
}

func readFileAsString(fn string) string {
	data, _ := filesystem.ReadString(fn)
	return data
}

type Widget interface {
	Name() string
	InitialInfo() Info
	OnClick(ClickEvent)
	Start()
}

type StatusBar struct {
	infos   []Info
	widgets map[string]Widget
	rwmutex sync.RWMutex
}

func NewStatusBar() *StatusBar {
	s := &StatusBar{}
	s.widgets = make(map[string]Widget, 0)
	s.infos = make([]Info, 0)
	fmt.Println("{\"version\": 1,\"click_events\": true}")
	fmt.Println("[")
	return s
}

func (s *StatusBar) Start() {
	for {
		clickDecoder := json.NewDecoder(os.Stdin)
		_, err := clickDecoder.Token()
		if err != nil {
			panic(err)
		}

		event := ClickEvent{}
		for clickDecoder.More() {
			fmt.Fprintf(os.Stdin, ",")
			err := clickDecoder.Decode(&event)
			if err != nil {
				continue
			}
			go s.sendClickEvent(event)
		}
	}
}

func (s *StatusBar) AddWidget(w Widget) {
	s.widgets[w.Name()] = w
	s.Add(w.InitialInfo())
	go w.Start()
}

func (s *StatusBar) sendClickEvent(ce ClickEvent) {
	s.widgets[ce.Name].OnClick(ce)
}

func (s *StatusBar) printInfo() {
	s.rwmutex.RLock()
	z, err := json.Marshal(s.infos)
	if err != nil {
		panic(err)
	}
	s.rwmutex.RUnlock()
	fmt.Printf("%s,\n", string(z))
}

func (s *StatusBar) Add(i Info) {
	if i.Markup == "pango" {
		i.Text = rewriteInfoText(i.Text)
	}
	position := 0
	var oldInfo Info
	found := false
	s.rwmutex.RLock()
	for currentNum, oldI := range s.infos {
		if oldI.Name == i.Name {
			oldInfo = oldI
			found = true
			position = currentNum
		}
	}
	s.rwmutex.RUnlock()
	s.rwmutex.Lock()
	if !found {
		s.infos = append(s.infos, i)
	} else {
		s.infos[position] = i
	}
	s.rwmutex.Unlock()
	if !reflect.DeepEqual(oldInfo, i) {
		s.printInfo()
	}

}
