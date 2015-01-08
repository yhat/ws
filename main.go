package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/websocket"
)

var (
	origin  string
	headers string
	version int
)

func init() {
	help = fmt.Sprintf(help, VERSION)
	flag.StringVar(&origin, "o", "http://0.0.0.0/", "websocket origin")
	flag.StringVar(&headers, "H", "", "a comma separated list of http headers")
	flag.IntVar(&version, "v", websocket.ProtocolVersionHybi13, "websocket version")
	flag.Parse()
}

const VERSION = "0.1"

var help = `ws - %s

Usage:
	ws [options] <url>

Use "ws --help" for help.
`

func parseHeaders(headers string) http.Header {
	h := http.Header{}
	for _, header := range strings.Split(headers, ",") {
		parts := strings.SplitN(header, ":", 2)
		if len(parts) != 2 {
			continue
		}
		h.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}
	return h
}

func main() {
	fatal := func(format string, a ...interface{}) {
		fmt.Fprintf(os.Stderr, format, a...)
		os.Exit(2)
	}
	target := flag.Arg(0)
	if target == "" {
		fatal(help)
	}
	config, err := websocket.NewConfig(target, origin)
	if err != nil {
		fatal("%s\n", err)
	}
	if headers != "" {
		config.Header = parseHeaders(headers)
	}
	config.Version = version
	ws, err := websocket.DialConfig(config)
	if err != nil {
		fatal("Error dialing %s: %v\n", target, err)
	}
	errc := make(chan error, 2)
	cp := func(dst io.Writer, src io.Reader) {
		_, err := io.Copy(dst, src)
		errc <- err
	}
	go cp(os.Stdout, ws)
	go cp(ws, os.Stdin)
	<-errc
}
