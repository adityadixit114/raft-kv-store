package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type NodeState string

const (
	Follower  NodeState = "Follower"
	Candidate NodeState = "Candidate"
	Leader    NodeState = "Leader"
)

type RaftNode struct {
	ID          string
	CurrentTerm int
	VotedFor    string
	State       NodeState
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		nodeID = "node-1"
	}

	node := &RaftNode{
		ID:          nodeID,
		CurrentTerm: 1,
		State:       Follower,
	}

	fmt.Printf("Starting Raft Node [%s] in state [%s] on port %s...\n", node.ID, node.State, port)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to bind port %s: %v", port, err)
	}
	defer listener.Close()

	fmt.Printf("Node %s successfully listening. Waiting for cluster heartbeat ping...\n", node.ID)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Printf("Shutting down Raft node %s cleanly.\n", node.ID)
}
