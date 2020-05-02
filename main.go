package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/nsf/termbox-go"
	"github.com/tomatosource/socklog"
)

const DEFAULT = termbox.ColorDefault

type Siv struct {
	RawFeed        []string
	Matches        []int
	InputChars     []rune
	CursorPosition int
}

func NewSiv() *Siv {
	termbox.SetCursor(0, 0)
	return &Siv{
		RawFeed:        []string{},
		Matches:        []int{},
		InputChars:     []rune{},
		CursorPosition: 0,
	}
}

func main() {
	socklogger := socklog.MustNew("localhost:8080")
	defer socklogger.Close()
	log.SetOutput(socklogger)

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)

	s := NewSiv()
	s.ReadStdIn()

	stay := true
	for stay {
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

func (s *Siv) ReadStdIn() {
	go func() {
		merged := io.MultiReader(os.Stdin, os.Stderr)
		scanner := bufio.NewScanner(merged)
		for scanner.Scan() {
			newLine := scanner.Text()
			s.RawFeed = append(s.RawFeed, newLine)
			if s.isMatch(newLine) {
				s.Matches = append(s.Matches, len(s.RawFeed)-1)
			}
			s.DrawFeed()
		}
	}()
}

func (s *Siv) isMatch(q string) bool {
	return true
}

func (s *Siv) DrawFeed() {
	from := 0
	if l := len(s.Matches); l > 5 {
		from = l - 5
	}
	for i, rawIdx := range s.Matches[from:] {
		SetRow(i+5, s.RawFeed[rawIdx])
	}
	termbox.Flush()
}
