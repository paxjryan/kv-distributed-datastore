// modified from cs426 lab1
// https://github.com/shixiao/cs426-spring2023-labs/tree/main/lab1

package server_lib

import (
	"context"
	"fmt"
	"log"
	"time"

	kvpb "mini-spanner/kv_service/proto"
	tspb "mini-spanner/time_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// options for service init
type KvServiceOptions struct {
	// Server address for the TimeService
	TimeServiceAddr string
}

type KvServiceServer struct {
	kvpb.UnimplementedKvServiceServer
	// options are read-only and intended to be immutable during the lifetime of the service
	options KvServiceOptions

	kvstore map[string]string
}

func MakeKvServiceServer(
	options KvServiceOptions,
) *KvServiceServer {
	log.Printf("Starting KvService")
	return &KvServiceServer{
		kvstore: make(map[string]string),
		options: options,
	}
}

// copied from frontend.go
func serviceConn(address string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.Dial(address, opts...)
}

func (db *KvServiceServer) GetVal(
	ctx context.Context,
	req *kvpb.GetKvRequest,
) (*kvpb.GetKvResponse, error) {

	// Error checks
	if len(req.Keys) == 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			"Requests should not be empty",
		)
	}

	vals := make([]*kvpb.KvPair, 0)

	var val string
	var ok bool

	for _, kv := range req.Keys {
		if val, ok = db.kvstore[kv.Key]; !ok {
			return nil, status.Error(
				codes.InvalidArgument,
				fmt.Sprintf(
					"KvService GetVal(): key %s not in kvstore",
					kv.Key,
				),
			)
		}
		log.Printf("server handling GetVal(%s) RPC: %s", kv.Key, val)
		vals = append(vals, &kvpb.KvPair{Key: kv.Key, Val: val + " at time " + db.getTimestamp().String()})
	}

	// When is time used?
	// Should we keep it?
	// t := db.getTimestamp()

	return &kvpb.GetKvResponse{Vals: vals}, nil
	// return &kvpb.GetKvResponse{Vals: []string{v}}, nil
}

func (db *KvServiceServer) SetVal(
	ctx context.Context,
	req *kvpb.SetKvRequest,
) (*kvpb.SetKvResponse, error) {

	// Assume all successes
	successes := make([]bool, 0)
	for _, kv := range req.KvPairs {
		db.kvstore[kv.Key] = kv.Val
		log.Printf("server handling SetVal(%s, %s) RPC", kv.Key, kv.Val)
		successes = append(successes, true)
	}

	return &kvpb.SetKvResponse{Successes: successes}, nil
}

func (db *KvServiceServer) getTimestamp() time.Time {
	timeServiceConn, err := serviceConn(db.options.TimeServiceAddr)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer timeServiceConn.Close()
	client := tspb.NewTimeServiceClient(timeServiceConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	out, err := client.GetTime(
		ctx,
		&tspb.GetTimeRequest{},
	)
	if err != nil {
		log.Fatalf(
			"Error getting time",
			err,
		)
	}

	log.Printf("%v", out.Time.AsTime())
	return out.Time.AsTime()
}
