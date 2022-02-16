package handler

import (
    "net/http" 

    "github.com/gorilla/mux"   
)

func InitRouter() *mux.Router {
    router := mux.NewRouter()

	// router.HandleFunc("/upload", uploadHandler).Methods("POST")
    router.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST") //Handle Function?
    
	router.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET")

    return router
}