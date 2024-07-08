package main

import (
	"context"
	"flag"
	"log"
	"time"

	kvpb "mini-spanner/kv_service/proto"
	tspb "mini-spanner/time_service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	kvServiceAddr = flag.String(
		"kv-service",
		"[::1]:8080",
		"The server address for the KvService in the format of host:port",
	)
	timeServiceAddr = flag.String(
		"time-service",
		"[::1]:8081",
		"server addr for TimeService",
	)
)

const NUMSHARDS = 9
const NUMSERVERS = 3

const READ = 0
const WRITE = 1

type Operation struct {
	Key string
	Val string
	Op  int
}

// var pool grpc.ConnPool
// var conns []*grpc.ClientConn

func serviceConn(address string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.Dial(address, opts...)
}

// func makeConnection(addr string) error {
// 	var opts []grpc.DialOption
// 	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	channel, err := grpc.Dial(addr, opts...)
// 	if err != nil {
// 		return err
// 	}
// 	defer channel.Close()

// 	kvClients = append(kvClients, kvpb.NewKvServiceClient(channel))

// 	return nil
// }

func sendSingleGetRequest(key string, client kvpb.KvServiceClient) (val string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kvp := kvpb.KvPair{Key: key}

	out, err := client.GetVal(
		ctx,
		&kvpb.GetKvRequest{Keys: []*kvpb.KvPair{&kvp}},
	)
	if err != nil {
		log.Fatalf(
			"Error retrieving value for key %s: %v\n",
			key,
			err,
		)
	}

	return out.Vals[0].Val
}

func sendSingleSetRequest(key string, val string, client kvpb.KvServiceClient) (success bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kvp := kvpb.KvPair{Key: key, Val: val}

	out, err := client.SetVal(
		ctx,
		&kvpb.SetKvRequest{KvPairs: []*kvpb.KvPair{&kvp}},
	)
	if err != nil {
		log.Fatalf(
			"Error setting value %s for key %s: %v\n",
			val,
			key,
			err,
		)
	}

	return out.Successes[0]
}

func sendSingleTimeRequest(client tspb.TimeServiceClient) (time time.Time) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	out, err := client.GetTime(
		context.Background(),
		&tspb.GetTimeRequest{},
	)
	if err != nil {
		log.Fatalf(
			"Error getting time",
			err,
		)
	}

	return out.Time.AsTime()
}

func processTransaction(ops []Operation) {
	for _, op := range ops {
		// determine shard
		// shard := util.GetShardForKey(op.Key, NUMSHARDS)

		if op.Op == READ { // read op

		} else { // write op

		}
	}
}

func main() {
	flag.Parse()

	log.Printf("kvServiceAddr: %s", *kvServiceAddr)

	// for i := 0; i < NUMSERVERS; i++ {
	conn, err := serviceConn(*kvServiceAddr)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	kvServiceClient := kvpb.NewKvServiceClient(conn)
	// }

	// pool = grpc.DialPool(context.Background(), option.WithGRPCConnectionPool(NUMSERVERS))

	log.Printf("connection to kvService successful")

	// log.Printf("timeServiceAddr: %s", *timeServiceAddr)

	// timeServiceConn, err := serviceConn(*timeServiceAddr)
	// if err != nil {
	// 	log.Fatalf("fail to dial: %v", err)
	// }
	// defer timeServiceConn.Close()
	// timeServiceClient := tspb.NewTimeServiceClient(timeServiceConn)

	// log.Printf("connection to timeService successful")

	var key, val string

	key = "hi"
	val = "abc"
	out := sendSingleSetRequest(key, val, kvServiceClient)
	log.Printf("set key %s val %s: success %v", key, val, out)

	key = "hi2"
	val = "abc2"
	out = sendSingleSetRequest(key, val, kvServiceClient)
	log.Printf("set key %s val %s: success %v", key, val, out)

	key = "hi2"
	val = sendSingleGetRequest(key, kvServiceClient)
	log.Printf("val for key %s is %s", key, val)

	key = "hi"
	val = sendSingleGetRequest(key, kvServiceClient)
	log.Printf("val for key %s is %s", key, val)

	// time := sendSingleTimeRequest(timeServiceClient)
	// log.Printf("%v", time)
}
