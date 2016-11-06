package note

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
    "encoding/json"
    "strconv"
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

func GetNote(id int) (*Note, error) {
    var n *Note
    noteId := strconv.Itoa(id)
    err := db.View(func (tx *bolt.Tx) error{
        var err error
        b := tx.Bucket([]byte("notes"))
        k := []byte(noteId)
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

func GetNotes() []Note {
    var notes []Note
    // var n Note
    db.View(func(tx *bolt.Tx) error {
        // var err error
        c := tx.Bucket([]byte("notes")).Cursor()
        for k, v := c.First(); k!=nil; k,v =c.Next() {
            fmt.Printf("all key=%s, value=%s\n", k, v)
            // err = json.Unmarshal(k, n)
            // if err != nil {
            //     return err
            // }
            // notes = append(notes, n)
        }
        return nil
    })
    return notes
}
