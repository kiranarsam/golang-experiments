package main

import (
  "fmt"
  "net"
  "net/rpc"
)

type Server struct {}

func (this *Server) Negate (i int64, reply *int64) error {
  *reply = -i
  return nil
}

func server() {
  rpc.Register(new(Server))
  ln, err := net.Listen("tcp", ":8030")
  if err != nil {
    fmt.Println(err)
    return
  }
  
  for {
    con, err := ln.Accept()
    if err != nil {
      continue
    }
    
    go rpc.ServeConn(con)
  }
}

func client() {
  con, err := rpc.Dial("tcp", "127.0.0.1:8030")
  if err != nil {
    fmt.Println(err)
    return
  }
  var result int64
  err = con.Call("Server.Negate", int64(999), &result)
  if err != nil {
    fmt.Println(err)
  } else { 
    fmt.Println("Server.Negate(999) = ", result)
  }
}

func main() {
  go server()
  go client()
  
  var input string
  fmt.Scanln(&input)
}

