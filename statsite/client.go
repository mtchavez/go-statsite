package statsite

import (
	"log"
	"net"
	"time"
)

type Client struct {
	Conn net.Conn
	addr string
}

func (c *Client) Emit(m Messenger) (bool, error) {
	return c.Emitter(m.String())
}

func (c *Client) Emitter(msg string) (ok bool, err error) {
	ok = false
	if c.Conn == nil {
		return
	}
	_, err = c.Conn.Write([]byte(msg))
	if err != nil {
		return
	}
	ok = true
	return
}

func NewClient(addr string) (client *Client, err error) {
	_, err = net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Println("Error resolving statsite: ", err)
		return
	}
	var conn net.Conn
	conn, err = net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		log.Println("Error connecting to statsite: ", err)
		return
	}
	client = &Client{Conn: conn, addr: addr}
	return
}
