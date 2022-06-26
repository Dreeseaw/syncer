package syncer

import (
    // "fmt"
    "log"
    // "context"

    // etcd & grpc
    "google.golang.org/grpc"
    "go.etcd.io/etcd/client/v3"
    resolver "go.etcd.io/etcd/client/v3/naming/resolver"
    // endpoints "go.etcd.io/etcd/client/v3/naming/endpoints"
    
    // "github.com/kelindar/column"
    // "github.com/kelindar/column/commit"
)

type Backend interface {
    
}

type GrpcBackend struct {
    Conn *grpc.ClientConn
}

func NewGrpcBackend(nodeId, proxyAddr string) *GrpcBackend {

    // connect to grpc-proxy
    log.Println("connecting to etcd client")
    cli, etcdErr := clientv3.NewFromURL(proxyAddr)
    if etcdErr != nil {
        panic(etcdErr)
        return nil
    }

    // create grpc broadcasting client
    res, resErr := resolver.NewBuilder(cli)
    if resErr != nil {
        panic(resErr)
    }

    conn, grpcErr := grpc.Dial("etcd:///syncer/service", grpc.WithResolvers(res))
    if grpcErr != nil {
        panic(grpcErr)
    }

    return &GrpcBackend{Conn: conn}
}

// --- Test Mock Version ---

type MockBackend struct {}

func NewMockBackend() *MockBackend {
    return &MockBackend{}
}
