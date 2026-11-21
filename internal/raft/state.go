package raft

type Role int 

const (
	Follower Role = iota
	Candidate
	Leader
)

// For Logs
func (r Role) String() string {
	switch r {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	default:
		return "Unknown!"
	}
}