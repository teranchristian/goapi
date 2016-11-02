package main

import (
    "fmt"
    "html"
    "net/http"
    "strings"
    "strconv"
    "./note"
)

var notesBucket = []byte("notes")


func getId(path string) (id int, error string) {
    parameter := strings.Split(path, "/")
    if (len(parameter) > 3){
        error = "endpoint does not"
    } else {
        idP, err := strconv.Atoi(parameter[2])
        if err != nil {
            fmt.Println(err)
        }
        id = idP
    }
    return id, error
}

// func (note *Note) save(db *bolt.DB) error {
//     // Store the user model in the user bucket using the username as the key.
//     err := db.Update(func(tx *bolt.Tx) error {
//         b, err := tx.CreateBucketIfNotExists(notesBucket)
//         if err != nil {
//             return err
//         }    

//         encoded, err := json.Marshal(note)
//         if err != nil {
//             return err
//         }
//         return b.Put([]byte(note.title), encoded)
//     })
//     return err
// }

func main() {
    note.Open()
    // note.Save()
    // type Note note.Note
    n := note.Note{Id:"123",Title:"title", Text:"text", Author:"author", Date:"date"}
    n.Save()
    n1, _ := note.GetNote("123433")
    fmt.Println(n1)
    http.HandleFunc("/notes/", func(w http.ResponseWriter, r *http.Request) {
        // if r.URL.Path != "/notes/" {
        //     http.NotFound(w, r)
        //     return
        // }

        id, error := getId(r.URL.Path)
        if len(error) > 0 {
            http.NotFound(w, r)
        }
        fmt.Println(id)
        fmt.Println(error)

        defer note.Close()

        switch r.Method {
            case "GET":
                fmt.Fprintf(w, "GET, %q", html.EscapeString(r.URL.Path))
            case "POST":
                fmt.Fprintf(w, "POST, %q", html.EscapeString(r.URL.Path))
            case "PUT":
                fmt.Fprintf(w, "PUT, %q", html.EscapeString(r.URL.Path))
            case "DELETE":
                fmt.Fprintf(w, "DELETE, %q", html.EscapeString(r.URL.Path))
            default:
                http.Error(w, "Invalid request method.", 405)
        }
    })

    http.ListenAndServe(":8080", nil)
}

// So the exercise is simple. I want you to write a small JSON API in Golang.
// It should support saving, editing, deleting and listing "notes", like for a simple notes app.
// Notes are just going to be "title", "text", "date created" and "author",
// and whatever else you feel is missing.
