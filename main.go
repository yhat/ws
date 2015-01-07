package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/websocket"
)

var origin string

func init() {
	flag.StringVar(&origin, "origin", "http://localhost/", "websocket origin")
	flag.Parse()
}

func main() {
	url := flag.Arg(0)
	if url == "" {
		fmt.Fprintf(os.Stderr, "Usage: ws <url>\n")
		os.Exit(2)
	}
	ws, err := websocket.Dial(url, "", "http://localhost/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error dialing %s: %v\n", url, err)
		os.Exit(2)
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
