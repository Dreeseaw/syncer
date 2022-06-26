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
    s := New("test", "localhost:2381")
    err := defaultTestColls(s)
    assert.Nil(t, err)
}

func TestCommitPipeline(t *testing.T) {
    s := New("test", "localhost:2381")
    err := defaultTestColls(s)

    // insrt dummy row
    s.Source.Insert(func (r column.Row) error {
        r.SetAny("id", "bob")
        r.SetAny("cnt", 2)
    })

    expected := <-s.stream
    actual := fromCommitPb(toCommitPb(expected))

    // compare comm & res
    assert.Equal(t, expected, actual)
}
