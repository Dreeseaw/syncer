package syncer

import (
    "github.com/kelindar/column"
    "github.com/kelindar/column/commit"
)

type Syncer struct {
    stream chan commit.Commit
    target *column.Collection 
}

func NewSyncer() *Syncer {
    return &Syncer{
        stream: make(chan commit.Commit),
        target: nil,
    }
}

// Append fulfills commit.Logger
func (s *Syncer) Append(comm commit.Commit) error {
    s.stream <- comm.Clone()
}

func (s *Syncer) Assign(coll *column.Collection) error {
    if s.target != nil {
        return errors.New("syncer already assigned")
    }
    s.target = coll
    return nil
}

func (s *Syncer) Start() {
    
    // start grpc client

    // start grpc server

    go func(){
        for change := s.stream {
            // send change to other nodes
        }
        return
    }()

}
