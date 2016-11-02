package note

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
    "encoding/json"
)

var db *bolt.DB

type Note struct {
    Id string
    Title string
    Text string
    Author string
    Date string
}

func Open() error {
    var err error
    db, err = bolt.Open("/tmp/test.db", 0644, nil);
    if err != nil {
        log.Fatal(err)
    }
    return err
}

func Close() {
    db.Close()
}

func (n *Note) Save() error {
    err := db.Update(func(tx *bolt.Tx) error {
        notes, err := tx.CreateBucketIfNotExists([]byte("notes"))
        if err != nil {
            return fmt.Errorf("Error Bucket %s", err)
        }
        enc, err := json.Marshal(n)
        if err != nil {
            return fmt.Errorf("error %s:%s", n.Id, err)
        }
        err = notes.Put([]byte(n.Id), enc)
        return err
    })
    return err
}

func GetNote(id string) (*Note, error) {
    var n *Note
    err := db.View(func (tx *bolt.Tx) error{
        var err error
        b := tx.Bucket([]byte("notes"))
        k := []byte(id)
        err = json.Unmarshal(b.Get(k), &n)
        if err != nil {
            return err
        }
        return nil
    })
    if err != nil {
        fmt.Printf("Could not get Note ID %s", id)
        return nil, err
    }
    return n, nil
}
