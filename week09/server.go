package week09

import (
	"fmt"
	"log"
	"net"

	"github.com/Terry-Mao/goim/api/protocol"
	"github.com/Terry-Mao/goim/pkg/bufio"
)

// Server ...
type Server struct {
	host string
	port string
}

// Client ...
type Client struct {
	conn net.Conn
}

// Config ...
type Config struct {
	Host string
	Port string
}

// New ...
func New(config *Config) *Server {
	return &Server{
		host: config.Host,
		port: config.Port,
	}
}

// Run ...
func (server *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn: conn,
		}
		go client.handleRequest()
	}
}

func (client *Client) handleRequest() {
	rd := bufio.NewReader(client.conn)
	for {
		p := protocol.Proto{}
		if err := p.ReadTCP(rd); err != nil {
			client.conn.Close()
			fmt.Printf("Message Read Err: %v\n", err)
			return
		}
		fmt.Printf("Message incoming: %+v\n", p)
		client.conn.Write([]byte(fmt.Sprintf("Message received: %v\n", string(p.Body))))
	}
}
