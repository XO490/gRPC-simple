package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-simple/internal/app/server"
	pb "grpc-simple/proto"
	"log"
	"net"
)

var (
	port = flag.Int("port", 9092, "the server port")
)

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	server.CheckErr(err, "")

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterProductServer(s, &server.Server{})

	log.Printf("Server start at %v", l.Addr())

	err = s.Serve(l)
	server.CheckErr(err, "failed to serve")
}
