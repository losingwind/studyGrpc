package main

import (
	"log"
	"net"
	"net/rpc"
	"studyGrpc/pkg"
)

type Server struct{}

func (s *Server) Server(request string, rsp *string) error {
	*rsp = "hello, client!"
	return nil
}

func main() {
	if err := pkg.Register(new(Server)); err != nil {
		log.Fatal(err.Error())
	}

	listener, err := net.Listen("tcp", pkg.Addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	rpc.Accept(listener)
}
