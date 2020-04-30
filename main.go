package main

import (
	"bufio"
	"os"
	"syscall"

	"github.com/rivo/tview"
)

func main() {
	// input := tview.NewInputField().
	// SetDoneFunc(func(key tcell.Key) {
	// // log.Println("input done")
	// })

	text := tview.NewTextView().SetWrap(true)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		// AddItem(input, 1, 1, true).
		AddItem(text, 0, 1, true)

	app := tview.NewApplication().SetRoot(flex, true)

	readStdIn(app, text)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func readStdIn(app *tview.Application, textView *tview.TextView) {
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		// set tview tty to stdin
		os.Stdin = os.NewFile(uintptr(syscall.Stderr), "/dev/tty")
		for scanner.Scan() {
			textView.Write([]byte("apples\n"))
		}
	}()
}
