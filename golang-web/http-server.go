package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

    // Sending response data in "res"
    fmt.Fprintf(res, "Welcome to my website!")
  })
  
  // static files to be include
  fs := http.FileServer(http.Dir("static/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  // Listen to the server
  http.ListenAndServe(":8080", nil)	
	
}

