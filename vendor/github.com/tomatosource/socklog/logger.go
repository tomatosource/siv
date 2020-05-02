package socklog

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Logger struct {
	conn *websocket.Conn
}

// Returns a pointer to a new socketlogger
// that will write to `addr`
func New(addr string) (*Logger, error) {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/log"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("dialing: %w", err)
	}
	return &Logger{conn: c}, nil
}

// Same as new, fatals on any error
func MustNew(addr string) *Logger {
	l, err := New(addr)
	if err != nil {
		log.Fatal(err)
	}
	return l
}

// Closes the websocket connection to the logging server
func (l *Logger) Close() {
	l.conn.Close()
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.conn.WriteMessage(websocket.TextMessage, p)
	return len(p), nil
}
