package main

import "github.com/nsf/termbox-go"

var eventQueue chan termbox.Event

func main() {
	termbox.Init()
	defer termbox.Close()

	eventQueue = make(chan termbox.Event, 100)

	go func () {
		for {
			eventQueue<-termbox.PollEvent()
		}
	} ()

	startPanel := NewStartPanel(eventQueue)
	startPanel.TakeControl()
}

func Print(x, y int, str string, fg, bg termbox.Attribute) {
	runes := []rune(str)
	for i, r := range runes {
		termbox.SetCell(x + i, y, r, fg, bg)
	}
}

