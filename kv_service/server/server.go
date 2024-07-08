// modified from cs426 lab1
// https://github.com/shixiao/cs426-spring2023-labs/tree/main/lab1

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "mini-spanner/kv_service/proto"
	usl "mini-spanner/kv_service/server_lib"

	"google.golang.org/grpc"
)

var (
	port            = flag.Int("port", 8080, "The server port")
	timeServiceAddr = flag.String(
		"time-service",
		"[::1]:8081",
		"Server address for the TimeService",
	)
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterKvServiceServer(
		s,
		usl.MakeKvServiceServer(usl.KvServiceOptions{
			TimeServiceAddr: *timeServiceAddr,
		}),
	)
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
