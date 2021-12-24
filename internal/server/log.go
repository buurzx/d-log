package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

//NewLog instantiating Log
func NewLog() *Log {
	return &Log{}
}

//Append record to existing slice and responds with record offset
func (l *Log) Append(record Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	record.Offset = uint64(len(l.records))
	l.records = append(l.records, record)
	return record.Offset, nil
}

//Read from log from a given offset
func (l *Log) Read(offset uint64) (Record, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if offset >= uint64(len(l.records)) {
		return Record{}, ErrorOffsetNotFound
	}

	return l.records[offset], nil
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrorOffsetNotFound = fmt.Errorf("offset not found")
