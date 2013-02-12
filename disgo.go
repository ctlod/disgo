package main

import "github.com/nsf/termbox-go"

func main() {
  // Must be done here or doesn't work at all!
  termbox.Init()
  defer termbox.Close()

	startPanel := NewStartPanel()
	startPanel.TakeControl()
}

