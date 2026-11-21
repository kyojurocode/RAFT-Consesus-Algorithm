package raft

import "time"

type Config struct {
	ID int
	Peers []int // ID of all other nodes
	ElectionTimeoutMin time.Duration
	ElectionTimeoutMax time.Duration
	HeartbeatInterval time.Duration
}

// Making a Default Config 
func DefaultConfig(id int, peers []int) Config {
	return Config{
		ID: id,
		Peers: peers,
		ElectionTimeoutMin: 150 * time.Millisecond,
		ElectionTimeoutMax: 300 * time.Millisecond,
		HeartbeatInterval: 50 * time.Millisecond,
	}
}