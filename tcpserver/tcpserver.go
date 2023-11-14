package tcpserver

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
)

type TCPServer struct {
	ctx       context.Context
	addr      string
	port      string
	listener  net.Listener
	isRunning bool
}

func NewTCPServer(ctx context.Context, addr string, port string) TCPServer {
	return TCPServer{
		ctx:  ctx,
		addr: addr,
		port: port,
	}
}

func (ts *TCPServer) Run() error {
	lst, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ts.addr, ts.port))
	if err != nil {
		return fmt.Errorf("start server error: %w", err)
	}
	ts.listener = lst

	log.Printf("start server on %s\n", ts.listener.Addr())

	e := make(chan error)
	go func() {
		for err := range e {
			log.Println(err)
		}
	}()

	ts.isRunning = true
	for ts.isRunning {
		conn, err := ts.listener.Accept()
		if err != nil {
			log.Printf("accept connection error: %s\n", err.Error())
			continue
		}

		fmt.Printf("connection accepted: %+v, remote addr: %s\n", conn, conn.RemoteAddr())

		go func() {
			e <- readFromConnect(conn)
		}()

	}
	return nil
}

func (ts *TCPServer) Stop() error {
	if ts.isRunning {
		log.Printf("stop server on %s\n", ts.listener.Addr())
		ts.isRunning = false
		return ts.listener.Close()
	}
	log.Printf("server %s is stopped already\n", ts.addr)
	return nil
}

func readFromConnect(conn net.Conn) error {
	buffer := make([]byte, 256)
	var message string
	for {
		_, err := conn.Write([]byte("> "))
		if err != nil {
			return fmt.Errorf("write connection error: %w", err)
		}
		n, err := conn.Read(buffer)
		if err == io.EOF {
			log.Printf("close connection %v\n", conn)
			return conn.Close()
		}
		if err != nil {
			return fmt.Errorf("read connection error: %w", err)
		}
		message = string(buffer[:n])
		log.Printf("conn: %v, recieved message: %s", conn, message)
	}
}
