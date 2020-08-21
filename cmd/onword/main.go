package main

import (
	"flag"

	"github.com/superloach/onword"
	"github.com/superloach/onword/trans/tcp"
)

var (
	tcpAddr = flag.String("tcp", "", "tcp address")
)

func main() {
	flag.Parse()

	o := onword.NewOnword()

	if *tcpAddr != "" {
		go tcp.NewTCP(o, *tcpAddr).Run()
	}

	select {}
}
