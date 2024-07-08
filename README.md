## Overview

Our team decided to implement a mini-version of Google’s Spanner in Go. We intended to base our project on lab 4’s sharded key-value store, and focus on the following components: single node serializability with 2-phase locking, cluster coordination with 2-phase commit and coordinator crash tolerance, and transaction timestamping with the TrueTime API.

Our design was inspired by the Google Spanner paper. For the data store model, instead of using a database with relational tables and a query language, we implement a key-value store that supports transactions involving multiple ‘set‘ and ‘get‘ calls. The transactions will be processed with 2-phase locking to guarantee serializability. Each node is a single in-memory key-value store that we assume to be reliable, and will host multiple shards. We will implement an atomic distributed commit using 2-phase commit, wherein timestamping will be implemented with our bare-bones version of the TrueTime API using the Network Time Protocol. For deployment, we get the “True Time” from a dedicated time server. For simplicity, the client will always serve as the transaction coordinator of 2-phase commit and is allowed to crash, and we will implement a simple crash-tolerance mechanism.


## Work division

Pax: adapted KV-store from lab4 and deployed it on Kubernetes cluster, implemented timestamping RPC, made demo video
Anna: implemented the transaction processing logic of the client, wrote part of the final writeup
Peter: implemented 2PL RPC and wrote framework code for 2PC, wrote part of the final writeup
