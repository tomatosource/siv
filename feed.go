package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/nsf/termbox-go"
)

type matcher func(string) bool

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

func strContains(s string) matcher {
	return func(line string) bool {
		return strings.Contains(line, s)
	}
}

func (s *Siv) DrawFeed() {
	_, h := termbox.Size()
	for i := h - 3; i >= 0; i-- {
		ClearRow(i)
		idx := len(s.Matches) - h + i
		if idx >= 0 && idx < len(s.Matches) {
			line := s.RawFeed[s.Matches[idx]]
			SetRow(i, line)
		}
	}
	termbox.Flush()
}

func (s *Siv) isMatch(line string) bool {
	q := string(s.InputChars)
	matchingFunc := strContains(q)
	if strings.HasPrefix(q, "/") {
		re, err := regexp.Compile(q[1:])
		if err != nil {
			return false
		}
		matchingFunc = func(line string) bool {
			return re.MatchString(line)
		}
	}
	return matchingFunc(line)
}

func (s *Siv) Refilter(q string) {
	matches := []int{}
	matchingFunc := strContains(q)
	if strings.HasPrefix(q, "/") {
		re, err := regexp.Compile(q[1:])
		if err != nil {
			s.SetBrokenRegex(true)
			return
		}
		s.SetBrokenRegex(false)
		matchingFunc = func(line string) bool {
			return re.MatchString(line)
		}
	}
	for i, line := range s.RawFeed {
		if matchingFunc(line) {
			matches = append(matches, i)
		}
	}
	s.Matches = matches
	s.DrawFeed()
}
