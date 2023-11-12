package tcpserver

import (
	"fmt"
	"log"
	"net"
)

type TCPServer struct {
	addr string
	port string
}

func NewTCPServer(addr string, port string) TCPServer {
	return TCPServer{
		addr: addr,
		port: port,
	}
}

func (ts TCPServer) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ts.addr, ts.port))
	if err != nil {
		return fmt.Errorf("start server error: %w", err)
	}
	defer listener.Close()

	log.Printf("start server addr: %s, port: %s\n", ts.addr, ts.port)

	e := make(chan error)
	go func() {
		for err := range e {
			log.Println("error:", err)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("accept connection error: %w", err)
		}
		fmt.Printf("connection accepted: %+v\n", conn)
		go readFromConnect(conn, e)
	}
}

func readFromConnect(conn net.Conn, e chan error) {
	defer func() {
		conn.Close()
		log.Printf("connection %+v was closed\n", conn)
	}()

	buffer := make([]byte, 256)
	var message string
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			e <- fmt.Errorf("read connection error: %w", err)
			return
		}
		message = string(buffer[:n])
		log.Printf("conn: %v, recieved message: %s", conn, message)
		_, err = conn.Write([]byte("ok\n"))
		if err != nil {
			e <- fmt.Errorf("write connection error: %w", err)
			return
		}
	}
}
