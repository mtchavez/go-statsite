package statsite

import (
	"log"
	"net"
	"time"
)

type Client struct {
	Conn *net.TCPConn
}

func (c *Client) Emit(m Messenger) (bool, error) {
	return c.Emitter(m.String())
}

func (c *Client) Emitter(msg string) (ok bool, err error) {
	ok = false
	_, err = c.Conn.Write([]byte(msg))
	if err != nil {
		return
	}
	ok = true
	return
}

func NewClient(addr string) (client *Client, err error) {
	var tcpaddr *net.TCPAddr
	tcpaddr, err = net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Println("Error resolving statsite: ", err)
		return
	}
	var conn *net.TCPConn
	net.DialTimeout("tcp", addr, time.Now().Add(1*time.Second))
	conn, err = net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		log.Println("Error connecting to statsite: ", err)
		return
	}
	client = &Client{Conn: conn}
	return
}
