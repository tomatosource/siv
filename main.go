package main

import (
	"os"

	"github.com/nsf/termbox-go"
)

const DEFAULT = termbox.ColorDefault

type Siv struct {
	RawFeed        []string
	Matches        []int
	InputChars     []rune
	CursorPosition int
	BrokenRegex    bool
}

func NewSiv() *Siv {
	return &Siv{
		RawFeed:        []string{},
		Matches:        []int{},
		InputChars:     []rune{},
		CursorPosition: 0,
		BrokenRegex:    false,
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer func() {
		termbox.Close()
	}()

	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse | termbox.InputEsc)

	s := NewSiv()
	s.DrawCursor()
	s.ReadStdIn()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlC {
				termbox.Close()
				os.Exit(0)
			}
			// if ev.Ch == 'c' && ev.Mod == termbox.Mod
			s.HandleKeyEvent(ev)
		// TODO case termbox.EventMouse:
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
