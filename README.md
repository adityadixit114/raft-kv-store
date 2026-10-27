# Raft KV Store

A distributed, fault-tolerant key-value store built in Go implementing the Raft Consensus Protocol from scratch.

## Highlights
- **Raft Consensus Protocol**: Leader election, randomized heartbeats, log replication, and commit management.
- **Strong Linearizability**: Consistent distributed read/write semantics across multiple nodes.
- **Write-Ahead Logging (WAL)**: Ensures total data durability across sudden crashes and machine restarts.
- **gRPC Communication**: Low-latency inter-node communication protocols (< 4ms p99 latency).

## Project Architecture
```
[Client Request] ---> [Raft Leader Node (gRPC)]
                            |
           +----------------+----------------+
           |                                 |
     [Follower Node 1]                 [Follower Node 2]
     (WAL File Storage)               (WAL File Storage)
```

## Running the Cluster
```bash
# Clone repository
git clone https://github.com/adityadixit114/raft-kv-store.git
cd raft-kv-store

# Run with Docker Compose
docker compose up --build
```
