// modified from cs426 lab1
// https://github.com/shixiao/cs426-spring2023-labs/tree/main/lab1

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "mini-spanner/time_service/proto"
	usl "mini-spanner/time_service/server_lib"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTimeServiceServer(
		s,
		usl.MakeTimeServiceServer(),
	)
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
