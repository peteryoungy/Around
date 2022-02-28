package model

type Post struct {
    Id      string `json:"id"`  // java jackson
    User    string `json:"user"`  // go language support
    Message string `json:"message"`  // `` vs "":  string id="json:\"id\""  string id = `json:"id"`
    Url     string `json:"url"`
    Type    string `json:"type"`
}


type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Age      int64  `json:"age"`
    Gender   string `json:"gender"`
}