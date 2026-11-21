package raft

import (
	"log/slog"
	"sync"
	"time"
)

type Node struct {
	mu sync.Mutex

	cfg Config
	transport Transport
	logger *slog.Logger

	// Persisten state (Figure 2) from paper
	currentTerm int 
	votedFor int
	log *raftLog
	
	// Volatile state
	role Role
	commitIndex int
	lastApplied int

	// Timers 
	electionDeadline time.Time
	stopCh chan struct{}
}

func NewNode(cfg Config, t Transport) *Node {
	return &Node {
		cfg: cfg,
		transport: t,
		logger: slog.Default().With("node", cfg.ID),
		votedFor: -1,
		log: newRaftLog(),
		role: Follower,
		stopCh: make(chan struct{}),
	}
}