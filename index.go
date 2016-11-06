package main

import (
    "fmt"
    "html"
    "net/http"
    "strings"
    "strconv"
    "./note"
    "errors"
    "encoding/json"
    "html/template"
)

var notesBucket = []byte("notes")


func getId(path string) (int, error) {
    var err error
    var id int
    notePath := strings.TrimPrefix(path, "/notes")
    notePath = strings.TrimLeft(notePath,"/")
    if len(notePath) == 0 {
        return id, err
    }

    parameter := strings.Split(notePath, "/")
    if (len(parameter) > 1){
         err = errors.New("endpoint does not")
         return id, err
    } else {
        if (len(parameter) == 1) {
            id, err = strconv.Atoi(parameter[0])
        }
    }
    return id, err
}

func main() {
    note.Open()
    defer note.Close()
    // note.Save()
    // type Note note.Note
    // n := note.Note{Id:"1",Title:"title", Text:"text", Author:"author", Date:"date"}
    // n1 := note.Note{Id:"123",Title:"title", Text:"text", Author:"author", Date:"date"}
    // n.Save()
    // n1.Save()

    http.HandleFunc("/notes/", func(w http.ResponseWriter, r *http.Request) {
        // if r.URL.Path != "/notes/" {
        //     http.NotFound(w, r)
        //     return
        // }
        

        switch r.Method {
            case "GET":
                id, err := getId(r.URL.Path)
                if err != nil {
                    fmt.Fprintf(w, "error %q", err)
                    return
                }
                
                if id != 0 {
                    n1, _ := note.GetNote(id)
                    t, _ := json.Marshal(n1)
                    fmt.Fprintf(w, "Response %q", template.JS(t))
                } else {
                    note.GetNotes();
                }
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
