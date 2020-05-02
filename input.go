package main

import (
	"unicode"

	"github.com/nsf/termbox-go"
)

func (s *Siv) HandleKeyEvent(e termbox.Event) {
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

	chars := make([]rune, len(s.InputChars)-1)
	if s.CursorPosition == len(s.InputChars) {
		chars = s.InputChars[:len(s.InputChars)-1]
	} else {
		for i := 0; i < s.CursorPosition-1; i++ {
			chars[i] = s.InputChars[i]
		}
		for i := s.CursorPosition; i < len(s.InputChars); i++ {
			chars[i-1] = s.InputChars[i]
		}
	}
	s.InputChars = chars
	s.MoveCursorLeft()
	s.DrawInput()
	s.Refilter(string(s.InputChars))
}

func (s *Siv) InsertRune(r rune) {
	chars := make([]rune, len(s.InputChars)+1)
	if s.CursorPosition == len(s.InputChars) {
		chars = append(s.InputChars, r)
	} else {
		for i := 0; i < s.CursorPosition; i++ {
			chars[i] = s.InputChars[i]
		}
		chars[s.CursorPosition] = r
		for i := s.CursorPosition; i < len(s.InputChars); i++ {
			chars[i+1] = s.InputChars[i]
		}
	}

	s.InputChars = chars
	s.MoveCursorRight()
	s.DrawInput()
	go s.Refilter(string(s.InputChars))
}

func (s *Siv) MoveCursorLeft() {
	if s.CursorPosition == 0 {
		return
	}
	s.CursorPosition--
	s.DrawCursor()
}

func (s *Siv) MoveCursorRight() {
	if s.CursorPosition >= len(s.InputChars) {
		return
	}
	s.CursorPosition++
	s.DrawCursor()
}

func (s *Siv) DrawCursor() {
	_, h := termbox.Size()
	termbox.SetCursor(s.CursorPosition, h-1)
}

func (s *Siv) SetBrokenRegex(status bool) {
	s.BrokenRegex = status
	s.DrawInput()
}

func (s *Siv) DrawInput() {
	w, h := termbox.Size()
	inputRow := h - 1
	ClearRow(inputRow)
	for i, c := range s.InputChars {
		termbox.SetCell(i, inputRow, c, DEFAULT, DEFAULT)
	}

	if s.CursorPosition < len(s.InputChars) {
		termbox.SetCell(
			s.CursorPosition, inputRow, s.InputChars[s.CursorPosition], DEFAULT, DEFAULT,
		)
	} else {
		termbox.SetCell(s.CursorPosition, inputRow, rune(' '), DEFAULT, DEFAULT)
	}

	if s.BrokenRegex {
		msg := "[INVALID REGEX] "
		msgLen := len(msg)
		for i, c := range msg {
			termbox.SetCell(w-msgLen+i, inputRow, c, termbox.ColorRed|termbox.AttrBold, DEFAULT)
		}
	}
}
