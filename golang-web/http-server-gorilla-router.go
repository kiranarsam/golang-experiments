package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func CreateBook(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res, "CreateBook ")
}

func ReadBook(res http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  fmt.Fprintf(res, "ReadBook : %s\n", vars["title"])
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res, "UpdateBook ")
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res, "DeleteBook ")
}


func main() {

  route := mux.NewRouter()
  
  route.HandleFunc("/books/{title}/page/{page}", func(res http.ResponseWriter, req *http.Request){
    vars := mux.Vars(req)
    title := vars["title"]
    page := vars["page"]
    
    fmt.Fprintf(res, "You've requested the book: %s on page %s\n", title, page)
  })
  
  route.HandleFunc("/books/{title}", CreateBook).Methods("POST")
  route.HandleFunc("/books/{title}", ReadBook).Methods("Get")
  route.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
  route.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
  
  http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

    // Sending response data in "res"
    fmt.Fprintf(res, "Welcome to my website!")
  })
  
  // static files to be include
  fs := http.FileServer(http.Dir("static/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  // Listen to the server
  http.ListenAndServe(":8080", route)	
	
}

