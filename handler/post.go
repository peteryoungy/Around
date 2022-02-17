package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"around/model"
	"around/service"

	"github.com/pborman/uuid"
)

var (
    mediaTypes = map[string]string{
        ".jpeg": "image",
        ".jpg":  "image",
        ".gif":  "image",
        ".png":  "image",
        ".mov":  "video",
        ".mp4":  "video",
        ".avi":  "video",
        ".flv":  "video",
        ".wmv":  "video",
    }
)

//?: why responseWriter is not a pointer?
func uploadHandler(w http.ResponseWriter, r *http.Request) {


    // Parse from body of request to get a json object.
    fmt.Println("Received one upload request")

    p := model.Post{
        Id: uuid.New(),
        User: r.FormValue("user"),
        Message: r.FormValue("message"),
    }

    file, header, err := r.FormFile("media_file")
    if err != nil {
        http.Error(w, "Media file is not available", http.StatusBadRequest)
        fmt.Printf("Media file is not available %v\n", err)
        return
    }

    suffix := filepath.Ext(header.Filename)
    if t, ok := mediaTypes[suffix]; ok {
        p.Type = t
    } else {
        p.Type = "unknown"
    }

    err = service.SavePost(&p, file)
    if err != nil {
        http.Error(w, "Failed to save post to backend", http.StatusInternalServerError)
        fmt.Printf("Failed to save post to backend %v\n", err)
        return
    }

	// Fprint: print to a specific position,not stdout
    fmt.Fprintf(w, "Post received: %s\n", p.Message)
}


func searchHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Println("Received one search request")

    user := r.URL.Query().Get("user")  // ?xxxx
    keywords := r.URL.Query().Get("keywords")

    var posts []model.Post
    var err error

    if user != "" {
        posts, err = service.SearchPostByUser(user)
    } else {

        // nothing: keywords = "" return all
        posts, err = service.SearchPostByKeywords(keywords)
    }
    
    // catch error
    if err != nil {
        // panic(err)  // too racial

        // return server error to frent end
        http.Error(w, "Failed to read data from backend", http.StatusInternalServerError)  // 500
        return
    }

    // return search result, array --- > json
    js, err := json.Marshal(posts)
    if err != nil {
        http.Error(w, "Failed to parse posts into JSON format", http.StatusInternalServerError)  // 500
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}