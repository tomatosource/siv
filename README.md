# siv

Live filtering of stdin

## Installation

`go get github.com/tomatosource/siv`

## Usage

Unfortunately there's a quirk in the TUI lib that needs stdin to remain open, if you have a non streaming input you might be able to `tail -f` or need to use other tools.

`./somethingWithOutput | siv`
`tail -f someLogFile.txt | siv`

### Query language

- If query string begins with `/` the remaining characters will be interpreted as a [regular expression](https://github.com/google/re2/wiki/Syntax).
