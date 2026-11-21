package raft

type Transport interface {
	RequestVote(to int, args *RequestVoteArgs) (*RequestVoteReply, error)
	AppendEntries(to int, args *AppendEntriesArgs) (*AppendEntriesReply, error)
}

