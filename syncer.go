package syncer

import (
    "fmt"
    "errors"

    "github.com/kelindar/column"
    "github.com/kelindar/column/commit"
)

type Syncer struct {
    stream chan commit.Commit
    target *column.Collection 
    backend Backend
}

func New(nodeId, proxyAddr string) *Syncer {
    var b Backend
    if nodeId == "test" {
        b = NewMockBackend()
    } else {
        b = NewGrpcBackend(nodeId, proxyAddr)
    }
    return &Syncer{
        stream: make(chan commit.Commit),
        target: nil,
        backend: b,
    }
}

// Append fulfills commit.Logger
func (s *Syncer) Append(comm commit.Commit) error {
    s.stream <- comm.Clone()
    return nil
}

func (s *Syncer) Assign(coll *column.Collection) error {
    if s.target != nil {
        return errors.New("syncer already assigned")
    }
    s.target = coll
    return nil
}

func (s *Syncer) Start() {

    go func(){
        for change := range s.stream {
            // if change was a replica, discard
            // else, send change to other nodes
            fmt.Println(change)
        }
        return
    }()

}
