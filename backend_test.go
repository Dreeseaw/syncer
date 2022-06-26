package syncer

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func Test_EtcdConnect(t *testing.T) {
    etcd_conn := NewGrpcBackend("test", "127.0.0.1:2381")
    assert.NotNil(t, etcd_conn)
}
