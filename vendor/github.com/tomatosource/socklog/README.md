# socklog

A basic websocket logging util library to help debug terminal UI applications, or any app thats tricky to write to log to std out, written in Go.

Built upon [Gorilla's Websocket package.]("github.com/gorilla/websocket")

## Installation

```
go get github.com/tomatosource/socklog
```

## Usage

### Start the server

```
$ socklog
Listening for logs...

```

### Initialise your logger
```go
socklogger := socklog.MustNew("localhost:8080")
defer socklogger.Close()
log.SetOutput(socklogger)
```

### Log as normal

```go
log.Print("hello world")
```

## Options

```
$ socklog --help
Usage of socklog:
  -addr string
    	http service address (default "localhost:8080")
```

## Other Features

None planned, but this whole repository is less than 100 lines of Go (not counting examples), so it should be fairly easy to poke around in and extend.
