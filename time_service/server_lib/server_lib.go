// modified from cs426 lab1
// https://github.com/shixiao/cs426-spring2023-labs/tree/main/lab1

package server_lib

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "mini-spanner/time_service/proto"
)

type TimeServiceServer struct {
	pb.UnimplementedTimeServiceServer
}

func MakeTimeServiceServer() *TimeServiceServer {
	log.Printf("Starting TimeService")
	return &TimeServiceServer{}
}

func (db *TimeServiceServer) GetTime(
	ctx context.Context,
	req *pb.GetTimeRequest,
) (*pb.GetTimeResponse, error) {
	log.Printf("time requested")
	return &pb.GetTimeResponse{Time: timestamppb.New(time.Now())}, nil
}
