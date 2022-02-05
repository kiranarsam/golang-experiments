package main

import (
 "fmt"
 "net"
 "encoding/gob"
)

func server() {
  // listen on a port
  ln, err := net.Listen("tcp", ":8000")
  if err != nil {
    fmt.Println(err)
    return
  }
  
  for {
    // accept a conncetion
    con, err := ln.Accept()
    if err != nil {
      fmt.Println(err)
      continue
    }
    
    // handle the conncetion
    go handleServerConnection(con)
  }
}

func handleServerConnection(con net.Conn) {
  // receive the mesage
  var msg string
  err := gob.NewDecoder(con).Decode(&msg)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Received", msg)
  }
  
  con.Close()
}

func client() {
  // connect to the server
  con, err := net.Dial("tcp", "127.0.0.1:8000")
  if err != nil {
    fmt.Println(err)
    return
  }
  
  // send the message
  msg := "Hello World"
  fmt.Println("Sending", msg)
  err = gob.NewEncoder(con).Encode(msg)
  if err != nil {
    fmt.Println(err)
  }
  
  con.Close()
}


func main() {
  go server()
  go client()
  
  var input string
  fmt.Scanln(&input)
}

