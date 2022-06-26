package syncer

import (
    "context"

    "github.com/kelindar/column/commit"
    pb "github.com/Dreeseaw/syncer/grpc"
)

type BackendServer struct {
    pb.UnimplementedBackendServer

    // stuff I want
    commChan chan commit.Commit
}

func NewBackendServer(cc chan commit.Commit) *BackendServer {
    return &BackendServer{
        commChan: cc,
    }
}

// Unary RPC
func (bs *BackendServer) Connect (ctx context.Context, nc *pb.NewClient) (*pb.Response, error) {
    // TODO
    return nil, nil
}

// Unary RPC (for now)
func (bs *BackendServer) Transfer (ctx context.Context, c *pb.Commit) (*pb.Response, error) {
    // TODO
    bs.commChan <- toCommit(c)
    return nil, nil
}

func fromCommitPb(inp *pb.Commit) commit.Commit {
    updates := inp.GetUpdates()
    ret := commit.Commit{
        ID: inp.GetID(),
        Chunk: commit.Chunk(inp.GetChunk()),
        Updates: make([]*commit.Buffer, len(updates)),
    }
    for ui, u := range updates {
        ret.Updates[ui] = &commit.Buffer{
            last: u.GetLast(),
            chunk: u.GetChunk(),
            buffer: u.GetBuffer(),
            chunks: u.GetChunks(),
            Column: u.GetColumn(),
        }
    }
    return ret
}

func toCommitPb(inp commit.Commit) *pb.Commit {
    pbBufs := make([]*pb.Buffer, len(inp.Updates))
    for ui, u := range inp.Updates {
        pbBufs[ui] = toBufferPb(u)
    } 
    return &pb.Commit{
        ID: inp.ID,
        Chunk: inp.Chunk,
        Updates: pbBufs,
    }
}

func toBufferPb(inp *commit.Buffer) *pb.Buffer {
    return &pb.Buffer{
        last: inp.last,
        chunk: inp.chunk,
        buffer: inp.buffer,
        chunks: inp.headers,
        Column: inp.Column,
    }
}
