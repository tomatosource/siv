package main

import (
	"log"
	"unicode"

	"github.com/nsf/termbox-go"
)

func (s *Siv) HandleKeyEvent(e termbox.Event) {
	log.Printf("%+v", e)
	if e.Key == termbox.KeyCtrlE {
		s.CursorPosition = len(s.InputChars)
		termbox.SetCursor(s.CursorPosition, 0)
	}
	if e.Key == termbox.KeyCtrlA {
		s.CursorPosition = 0
		termbox.SetCursor(s.CursorPosition, 0)
	}
	if e.Key == termbox.KeyArrowLeft {
		s.MoveCursorLeft()
	}
	if e.Key == termbox.KeyArrowRight {
		s.MoveCursorRight()
	}

	if e.Key == termbox.KeyBackspace || e.Key == termbox.KeyBackspace2 {
		s.HandleBackspace()
	}
	if unicode.IsGraphic(e.Ch) {
		s.InsertRune(e.Ch)
	}
	if e.Key == termbox.KeySpace {
		s.InsertRune(rune(' '))
	}
}

func (s *Siv) HandleBackspace() {
	if s.CursorPosition == 0 {
		return
	}

	if s.CursorPosition == len(s.InputChars) {
		s.InputChars = s.InputChars[:len(s.InputChars)-1]
		s.MoveCursorLeft()
		s.DrawInput()
		return
	}

	chars := make([]rune, len(s.InputChars)-1)
	for i := 0; i < s.CursorPosition-1; i++ {
		chars[i] = s.InputChars[i]
	}
	for i := s.CursorPosition; i < len(s.InputChars); i++ {
		chars[i-1] = s.InputChars[i]
	}
	s.InputChars = chars
	s.MoveCursorLeft()
	s.DrawInput()

}

func (s *Siv) InsertRune(r rune) {
	if s.CursorPosition == len(s.InputChars) {
		s.InputChars = append(s.InputChars, r)
		s.MoveCursorRight()
		s.DrawInput()
		return
	}

	chars := make([]rune, len(s.InputChars)+1)
	for i := 0; i < s.CursorPosition; i++ {
		chars[i] = s.InputChars[i]
	}
	chars[s.CursorPosition] = r
	for i := s.CursorPosition; i < len(s.InputChars); i++ {
		chars[i+1] = s.InputChars[i]
	}

	s.InputChars = chars
	s.MoveCursorRight()
	s.DrawInput()
}

func (s *Siv) MoveCursorLeft() {
	if s.CursorPosition == 0 {
		return
	}
	s.CursorPosition--
	termbox.SetCursor(s.CursorPosition, 0)
}

func (s *Siv) MoveCursorRight() {
	if s.CursorPosition >= len(s.InputChars) {
		return
	}
	s.CursorPosition++
	termbox.SetCursor(s.CursorPosition, 0)
}

func (s *Siv) DrawInput() {
	ClearRow(0)
	for i, c := range s.InputChars {
		log.Printf("%v\n", c)
		termbox.SetCell(i, 0, c, DEFAULT, DEFAULT)
	}

	if s.CursorPosition < len(s.InputChars) {
		termbox.SetCell(s.CursorPosition, 0, s.InputChars[s.CursorPosition], DEFAULT, DEFAULT)
	} else {
		termbox.SetCell(s.CursorPosition, 0, rune(' '), DEFAULT, DEFAULT)
	}
	log.Printf("%s\n\n", string(s.InputChars))
}
