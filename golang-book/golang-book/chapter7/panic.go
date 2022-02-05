package main

import "fmt"

func first() {
  fmt.Println("1st")
}

func second() {
  fmt.Println("2nd")
}

func main() {

  // sceanrio 1
  defer second()
  first()
 
  fmt.Println("End of the main func")

  // sceanrio 2 (Not working properly as expected) check: example: panic02.go file  
  panic("PANIC")
  str := recover()
  fmt.Println(str)
}
