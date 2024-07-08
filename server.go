// NOTE: implementation of this pseudocode will eventually go in kv_service/server_lib/server_lib.go
// Just need to finish testing basic RPC implementation first

import (
	"fmt"
    "time"
)

// other rpcs:
//  getRead(txn) handler - return requested data
//  setRead(txn) handler - R lock requested shards, return requested data
//  extractPromise(txn) handler - W lock requested shards + prepare write (promise granted) or abort (promise denied)
//	globalCommit(txn id) handler - perform write + R/W unlock affected shards
//  gossip(txn id) handler - respond to peer request 


// pseudocode for server
type SpannerServer struct {
	NumShards	int
	Shards 		[]...
	// each shard should have its own value, 
	// which is an integer of value 0, 1, or 2 (Free, RL, WL),
	// and a local RWMutex that locks the value
}	

type SpannerServerOptions struct {
	// Based on KV value store config
}

const FREE = 0
const READ = 1
const WRITE = 2

func MakeSpannerServer(options SpannerServerOptions)(*SpannerServer, error){
	// Initialize Shards
}



/* 
type Operation struct {
	Op		string
	Key		string
	Val		string
}

type Transaction struct {
	Id		int			
	Ops 	[]*Operation	
}

type AcquireShardLockRequest struct {
	Ops		[]*Operation
}

type AcquireShardLockResponse struct {
	...
	Success bool
}

type FreeShardLockRequest struct {
	Ops 	[]*Operation
}

type FreeShardLockResponse struct {
	...
	Success bool
}
*/


func (server *SpannerServer) AcquireShardLock(
	ctx context.Context,
	req *AcquireLockRequest
) (*AcquireLockResponse, error) {

	// construct op table for each shard
	shardIdMap := make(map[int]int)
	for _, op := range req.Ops {
		shardId := shard_hash_function(op.Key)
		var mode int
		if op.Op == "get" {
			mode = READ
		} else if op.Op == "set" {
			mode = WRITE
		} else {
			return nil, status.Error(
				codes.Unimplemented,
				"Invalid operation"
			)
		}

		if val, ok := shardIdMap[shardId]; ok {
			shardIdMap[shardId] = Max(shardIdMap[shardId], mode)
		} else {
			shardIdMap[shardId] = mode
		}
	}

	var success bool
	// try 5 times
	for i := 0; i < 5; i++ {
		for shardId, mode := range shardIds {
			if (server.Shards[shardId].lk.TryLock()){ 
				server.Shards[shardId].lk.Lock()
				if mode == WRITE {
					success = (server.Shards[shardId].lkval == FREE)
					if success {
						server.Shards[shardId].lkval = mode
					} else {
						break
					}
				} else if mode == READ {
					success = (server.Shards[shardId].lkval <= READ)
					if success {
						server.Shards[shardId].lkval = mode
						server.Shards[shardId].readcnt++
					} else {
						break
					}
				} 
				server.Shards[shardId].lk.Unlock()
			} else {
				success = false
			}
		}
		if success {
			break
		} else {
			time.Sleep(10 * time.Milisecond)
		}
	}
	resp := &AcquiredShardLockResponse{Success: success}
	return resp, nil
}


func (server *SpannerServer) FreeShardLock(
	ctx context.Context,
	req *FreeShardLockRequest
) (*FreeShardLockResponse, error) {
	
	// construct op table for each shard, but used for freeing
	shardIdMap := make(map[int]int)
	for _, op := range req.Ops {
		shardId := shard_hash_function(op.Key)
		var mode int
		if op.Op == "get" {
			mode = READ
		} else if op.Op == "set" {
			mode = WRITE
		} else {
			return nil, status.Error(
				codes.Unimplemented,
				"Invalid operation"
			)
		}

		if val, ok := shardIdMap[shardId]; ok {
			shardIdMap[shardId] = Max(shardIdMap[shardId], mode)
		} else {
			shardIdMap[shardId] = mode
		}
	}

	success := true
	for shardId, mode := range shardIds {
		server.Shards[shardId].lk.Lock()
		if mode != server.Shards[shardId].lkval {
			success = false
		} else if mode == WRITE {
			server.Shards[shardId].lkval = FREE
		} else if mode == READ {
			server.Shards[shardId].readcnt--
			if server.Shards[shardId].readcnt < 0 {
				success = false
			} else if server.Shards[shardId].readcnt == 0 {
				server.Shards[shardId].lkval = FREE
			}
		} 
		server.Shards[shardId].lk.Unlock()
	}

	resp := &FreeShardLockResponse{Success: success}
	return resp, nil
}

// releaseShardLock can be a wrapper on the client side

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}