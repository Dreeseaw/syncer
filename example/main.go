package main

import (
    "fmt"
    "time"

    "github.com/Dreeseaw/syncer"
    "github.com/kelindar/column"
)

func main() {
    fmt.Println("syncer trying to connect")
    s := syncer.New("myNodeId", "grpc-proxy:2381")
    fmt.Println("created syncer")

    myColl := column.NewCollection(column.Options{
        Writer: s,
    })

    if err := s.Assign(myColl); err != nil {
        panic(err)
    }

    myColl.CreateColumn("id", column.ForString())
    myColl.CreateColumn("cnt", column.ForInt())

    // wait for grpc-proxy service discovery
    time.Sleep(2 * time.Second)

    s.Start()

    // print out all other nodes
    // TODO

    // do an insert on each node
    myColl.Insert(func (r column.Row) error {
        r.SetString("id", "bob")
        r.SetInt("cnt", 6)
        return nil
    })

    // wait for sync to kick in
    time.Sleep(2 * time.Second)

    // print out entire collection
    myColl.Query(func (txn *column.Txn) error {
        return txn.Range(func (i uint32) {
            rowObj := make(column.Object)
            id, _ := txn.String("id").Get()
            cnt, _ := txn.Int("cnt").Get()
            rowObj["id"] = id
            rowObj["cnt"] = cnt
            fmt.Printf("%v: %v\n", i, rowObj)
        })
    })
    return
}
