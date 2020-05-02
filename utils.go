package main

import "github.com/nsf/termbox-go"

func ClearRow(row int) {
	w, _ := termbox.Size()
	for i := 0; i < w; i++ {
		termbox.SetCell(i, row, rune(' '), DEFAULT, DEFAULT)
	}
}

func SetRow(row int, msg string) {
	ClearRow(row)
	for i, c := range msg {
		termbox.SetCell(i, row, c, DEFAULT, DEFAULT)
	}
}
