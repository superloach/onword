package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/superloach/onword"
)

type TCP struct {
	*onword.Onword

	Address string
}

func NewTCP(o *onword.Onword, address string) *TCP {
	return &TCP{
		Onword:  o,
		Address: address,
	}
}

func (t *TCP) Run() error {
	ln, err := net.Listen("tcp", t.Address)
	if err != nil {
		err = fmt.Errorf("ln tcp %q: %w", t.Address, err)
		return err
	}

	println("listening tcp", ln.Addr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			err = fmt.Errorf("accept conn: %w", err)
			return err
		}

		go t.Handle(conn)
	}
}

func (t *TCP) Handle(conn net.Conn) {
	ses := t.Open()
	defer ses.Close()

	buf := bufio.NewReader(conn)

	go func() {
		for {
			bline, err := buf.ReadBytes('\n')
			if err != nil {
				println(err.Error())
				conn.Close()
				return
			}

			line := strings.TrimRight(string(bline), "\r\n")

			fmt.Println("read", line)

			ses.Input <- line
		}
	}()

	go func() {
		for {
			line := <-ses.Output

			fmt.Println("write", line)

			_, err := conn.Write([]byte(line + "\n"))
			if err != nil {
				println(err.Error())
				conn.Close()
				return
			}
		}
	}()

	ses.Loop()
}
