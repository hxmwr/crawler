package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A int
	B int
}

type Service interface {
	add( b int) int
}

type Rpcsrv struct {}

func (r *Rpcsrv) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	s := Rpcsrv{}
	err := rpc.Register(&s)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := listener.Accept()
		jsonrpc.ServeConn(conn)
	}
}
