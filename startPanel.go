package main

import "github.com/nsf/termbox-go"
import "time"

type StartPanel struct {
	AnyKeyColor termbox.Attribute
	GuiTicker *time.Ticker
	TitleStr, AttrStr, AckStr string
	Input chan termbox.Event
}

func NewStartPanel () *StartPanel {
	sP := new(StartPanel)
  sP.AnyKeyColor = termbox.ColorWhite
	sP.GuiTicker = time.NewTicker(500 * time.Millisecond)
	sP.TitleStr = "Dwarves in Space!"
	sP.AttrStr = "A tribute to Dwarf Fortress"
	sP.AckStr = "Dedicated to Emma, for her patience."
	sP.Input = make(chan termbox.Event, 100)

  go func () {
    for {
      sP.Input<-termbox.PollEvent()
    }
  } ()

	return sP
}

func (sP* StartPanel) TakeControl () {
	sP.Clear()
	sP.Draw()
	termbox.Flush()
	for {
		guiTick := 0
		var e termbox.Event
		select {
		case e = <-sP.Input:
		case <-sP.GuiTicker.C:
			guiTick = 1
		}
		if (guiTick == 1 || e.Type == termbox.EventResize) {
			sP.Clear()
			sP.Draw()
			termbox.Flush()
		} else {
			sP.GuiTicker.Stop()
			break
		}
	}
}

func (sP* StartPanel) Clear () {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
}

func (sP* StartPanel) Draw () {
	width, height := termbox.Size()
	Print(width/2 - len(sP.TitleStr) / 2, height / 4, sP.TitleStr, termbox.ColorWhite | termbox.AttrBold , termbox.ColorBlack)
  Print(width/2 - len(sP.AttrStr) / 2, height / 4 + 1, sP.AttrStr, termbox.ColorWhite, termbox.ColorBlack)
  Print(width/2 - len(sP.AckStr) / 2, height - 1, sP.AckStr, termbox.ColorBlue, termbox.ColorBlack)
  Print(width/2 - len("Press Any Key") / 2, 2 * height / 3, "Press Any Key", sP.AnyKeyColor, termbox.ColorBlack)
	if sP.AnyKeyColor == termbox.ColorWhite {
		sP.AnyKeyColor = termbox.ColorBlack
	} else {
		sP.AnyKeyColor = termbox.ColorWhite
	}
}

func Print(x, y int, str string, fg, bg termbox.Attribute) {
  runes := []rune(str)
  for i, r := range runes {
    termbox.SetCell(x + i, y, r, fg, bg)
  }
}

