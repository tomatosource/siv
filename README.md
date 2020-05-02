# siv

Live filtering of stdin

## Installation

`go get github.com/tomatosource/siv`

## Usage

`./somethingWithOutput | siv`
`tail -f someLogFile.txt | siv`

### Query language

- If query string begins with `/` the remaining characters will be interpreted as a [regular expression](https://github.com/google/re2/wiki/Syntax).
