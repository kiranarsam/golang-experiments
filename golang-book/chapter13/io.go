package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("test.txt")
  if err != nil {
    // handle the error here
    fmt.Println("Error opening a file")
    return
  }
  defer file.Close()
  
  stat, err := file.Stat()
  if err != nil {
    return
  }
  
  //read the file
  bs := make([]byte, stat.Size())
  _, err = file.Read(bs)
  if err != nil {
    return
  }
  
  
  str := string(bs)
  fmt.Println(str)
}

