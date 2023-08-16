// NOTE: implementation of this pseudocode will eventually go in frontend/frontend.go
import (
	"context"
	"time"
)


// pseudocode for client / coordinator
type SpannerClient struct {
	NumServers 		int
	ServerConns 	[]*grpc.Conn
	ShardServerMap	map[string]int
}

type Operation struct {
	Op		string
	Key		string
	Val		string
}

type Transaction struct {
	Id		int			
	Ops 	[]*Operation	
}

type CommitTable struct {
	CommitId	int
	Commits		map[string]string
}

// Spanner Client should be able to call AcquireShardLock RPC

func (client *SpannerClient) ProcessTransaction(
	txn *Transaction
)(*CommitTable, error){
	// split transaction into operations grouped by server
	ops_grouped_by_server := make(map[int][]*Operation)
	for _, op := txn.Ops {
		server_id = server_hash_function(op.Key)
		ops_grouped_by_server[server_id] = append(ops_grouped_by_server[server_id], op)
	}


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 2PL: Acquire lock
	acquire_lock_phase_success := false
	acquired_server_ids := make(map[int]bool) // set of server ids where locks have been acquired
	num_retry := 2
	num_trials := 0
	while (!acquire_lock_phase_success) && num_trials < num_retry {
		for server_id, ops := range ops_grouped_by_server{
			req := &proto.AcquireShardLockRequest {Ops: ops}
			resp, _ := client.AcquireShardLock(ctx, req)
			if resp.Success{
				acquired_server_ids[server_id] = true
				acquire_lock_phase_success = true			
			} else {
				// abort
				acquire_lock_phase_success = false
				// free all acquired locks on abort
				for server_id, _ := range acquired_server_ids {
					req := &proto.FreeShardLockRequest {Ops: ops_grouped_by_server[server_id]}
					resp, _ := client.FreeShardLock(ctx, req)
					if !resp.Success {
						// Debug, since free lock shouldn't fail
					}
				}
				break
			}
		}
		time.Sleep(100 * time.Milisecond) // try again
		num_trials++
	}

	if !acquire_lock_phase_success {
		return nil, status.Error(
			codes.Aborted,
			"Can't acquire lock",
		)
	}
	
	// Create commit table by Sending Get Requests:
	// create client, establish connection, call Get
	// the get calls are already "within locks", no need to get lock on server side
	commitTable = &CommitTable{CommitId: txn.Id, Commits: make(map[string]string)}
	getCh := make(chan map[string][string], len(ops_grouped_by_server)) // type of channel should be return type of batched GetVal function

	var wg sync.WaitGroup
	for server_id, ops := range ops_grouped_by_server {
		// create GetVal request based on ops
		wg.Add(1)
		kv := make([]*proto.KvPair, 0)
		for _, op := range ops {
			kv = append(kv, &proto.KvPair{Key: op.Key, Val: op.Val})
		}
		req := &proto.GetKvRequest{Keys: kv}
		go func(ch chan map[string][string]){
			resp := client.GetVal(...)
			defer wg.Done()
			ch <- resp.Vals
		}(getCh)
	}
	
	// Wait and Close channel
	wg.Wait()	
	close(getCh)

	// Aggregate maps
	// Assume the keys do not overlap
	for kvMap := range getCh {
		for key, val := range kvMap {
			commitTable.Commits[key] = val
		}
	}

	// Process actual transaction

	// 1. Go through each operation in the transaction
	// Split operation by server
	setOps_by_server := make(map[int][]*Operation)
	for _, op := txn.Ops {
		if op.Op == "set" {
			// Update ops val
			server_id = server_hash_function(op.Key)
			commitTable.commits[op.Key] = op.Val
			setOps_by_server[server_id] = append(setOps_by_server, op)
		}
	}

	// 2. 2PC: Send TrySet Request
	setResponseCh := make(chan []bool, len(setOps_by_server)) // type of channel should be return type of batched SetVal function

	for server_id, ops := range setOps_by_server {
		// create GetVal request based on ops
		wg.Add(1)
		kv := make([]*proto.KvPair, 0)
		for _, op := range ops {
			kv = append(kv, &proto.KvPair{Key: op.Key, Val: commitTable.commits[op.Key]})
		}
		req := &proto.SetKvRequest{KvPairs: kv}
		go func(ch chan []bool){
			resp := client.TrySetVal(...)
			defer wg.Done()
			ch <- resp.Vals
		}(setResponseCh)
	}

	wg.Wait()
	close(setResponseCh)

	setError := false

	// 3. Check set request responses (successes)
	// Return error on error?
	for successes := range setResponseCh {
		for _, success := successes {
			if !success {
				setError = true
			}
		}
	}

	// 4. 2PC: send set requests
	if !setError {
		for server_id, ops := range setOps_by_server {
			// create GetVal request based on ops
			wg.Add(1)
			kv := make([]*proto.KvPair, 0)
			for _, op := range ops {
				kv = append(kv, &proto.KvPair{Key: op.Key, Val: commitTable.commits[op.Key]})
			}
			req := &proto.SetKvRequest{KvPairs: kv}
			go func(ch chan []bool){
				resp := client.SetVal(...)
				defer wg.Done()
				ch <- resp.Vals
			}(setResponseCh)
		}
	
		wg.Wait()
		close(setResponseCh)
	} 

	// 2PL: Free lock
	for server_id, ops := range ops_grouped_by_server{
		req := &proto.FreeShardLockRequest {Ops: ops}
		resp, _ := client.FreeShardLock(ctx, req)
		// check error
	}
	return commitTable, setError
}
