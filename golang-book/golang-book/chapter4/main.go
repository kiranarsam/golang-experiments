package main

import "fmt"

func main() {

  var x string = "Hello World"
  fmt.Println(x)
  
  var y string
  y = "first"
  fmt.Println(y)
  y = "second"
  fmt.Println(y)
  
  var z string
  z = "third "
  z = z + "fourth"
  fmt.Println(z)
  z += " fifth"
  fmt.Println(z)
  
  fmt.Println("x == y", (x == y))
  
  // Compiler will treat a new variable
  value := 5
  fmt.Println(value)
  
  
}

