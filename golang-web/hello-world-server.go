package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

    // Sending response data in "res"
    fmt.Fprintf(res, "Hello, you've requested: %s\n", req.URL.Path)
  })

  // Listen to the server
  http.ListenAndServe(":8080", nil)	
	
}

