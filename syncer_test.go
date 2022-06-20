package syncer

import (
    "testing"

    "github.com/kelindar/column"
    "github.com/stretchr/testify/assert"
)


func defaultTestColls(s *Syncer) error {
    source := column.NewCollection(column.Options{
        Writer: s,
    })
    if err := s.Assign(source); err != nil {
        return err
    }
    source.CreateColumn("id", column.ForString())
    source.CreateColumn("cnt", column.ForInt())
    return nil
}

func TestSyncer(t *testing.T) {
    s := NewSyncer()
    err := defaultTestColls(s)
    assert.Nil(t, err)
}
