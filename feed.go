package main

import (
	"bufio"
	"io"
	"os"

	"github.com/nsf/termbox-go"
)

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
