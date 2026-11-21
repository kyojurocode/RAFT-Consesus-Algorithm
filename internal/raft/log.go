package raft

type LogEntry struct {
	Term int 
	Command []byte
}

type raftLog struct {
	entries []LogEntry
}

func newRaftLog() *raftLog {
	return &raftLog{entries: []LogEntry{{Term: 0}}}
}

func (l *raftLog) lastIndex() int {
	return len(l.entries) - 1
}

func (l *raftLog) lastTerm() int {
	return l.entries[len(l.entries) - 1].Term
}

func (l *raftLog) termAt(i int) int {
	return l.entries[i].Term
}

func (l *raftLog) append(e LogEntry) {
	l.entries = append(l.entries, e)
}