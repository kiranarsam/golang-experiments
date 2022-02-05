package main

import (
  "fmt"
  "time"
  "math/rand"
)

func dump(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}

func main() {
  for i:=0; i < 10; i++ {
    go dump(i)
  }
  //dump(1)
  var input string
  fmt.Scanln(&input)
}

