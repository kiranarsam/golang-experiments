package main

import "fmt"

func main() {

  i := 1
  for i <= 10 {
    fmt.Println(i)
    i = i + 1
  }
  
  for x := 21; x <= 30; x++ {
    if x % 2 == 0 {
      fmt.Println(x, " even")
    } else {
      fmt.Println(x, " odd")
    }
  }
  
  y := 10
  if y > 10 {
    fmt.Println("Big")
  } else { 
    fmt.Println("Small")
  }
  
  y = 10
  switch y {
    case 10: fmt.Println("Ten")
    case 20: fmt.Println("Twenty")
    case 30: fmt.Println("Thirty")
    default: fmt.Println("default")
  }
}

