package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

func main() {

	n := NewNode("127.0.0.1:10001")
	fmt.Println("new node created")
	time.Sleep(30 * time.Second)

	fmt.Println("Killing it")
	close(n.StopChan)

	time.Sleep(3 * time.Second)
	fmt.Println("Exiting without problems")
}

type Node struct {
	StopChan chan struct{}
}

func (n *Node) DoNothing(_ *struct{}, reply *int) error {
        *reply = 5
	return nil
}

func NewNode(addr string) *Node {

	n := Node{make(chan struct{})}
	server := NewServer()
	if err := server.Register(&n); err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go server.Serve(listener)
	go func() {
		select {
		case <-n.StopChan:
			listener.Close()
		}
		fmt.Println("close listener")
	}()
	return &n
}

//

type Server struct {
	*rpc.Server
}

func (srv Server) Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go srv.ServeConn(conn)
	}
}

func NewServer() *Server {
	return &Server{Server: rpc.NewServer()}
}
