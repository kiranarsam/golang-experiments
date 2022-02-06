package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc


func Logging() Middleware {
  // create a new Middleware
  return func(f http.HandlerFunc) http.HandlerFunc {
    // define the http.HandlerFunc
    return func(w http.ResponseWriter, r *http.Request) {
      // Do middleware things
      start := time.Now()
      defer func() {
        log.Println(r.URL.Path, time.Since(start))
      }()
      
      // call the next middleware/handler in chain
      f(w,r)
    }
  }
}

func Method(m string) Middleware {
  // create a new middleware
  return func(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      // Do middleware things
      if r.Method != m {
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
      }
      
      // call the next middleware/handler in chain
      f(w, r)
    }
  }
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
  for _, m := range middlewares {
    f = m(f)
  }
  
  return f
}


func Hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello World from advanced Middleware")
}

func main() {
  http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
  http.ListenAndServe(":8080", nil)
  
}
