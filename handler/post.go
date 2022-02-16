package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"around/model"
	"around/service"
)

//?: why responseWriter is not a pointer?
func uploadHandler(w http.ResponseWriter, r *http.Request) {


    // Parse from body of request to get a json object.
    fmt.Println("Received one upload request")
    decoder := json.NewDecoder(r.Body)
    var p model.Post
    if err := decoder.Decode(&p); err != nil {
        panic(err)
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