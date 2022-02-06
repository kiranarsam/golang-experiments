package main 

import (
  "fmt"
  "math"
)

const CONSTANT string = "constant"

func main() {
  fmt.Println(CONSTANT)
  
  const NUMBER = 500000000
  
  const DECIMAL = 3e20 / NUMBER
  fmt.Println(DECIMAL)
  
  fmt.Println(int64(DECIMAL))
  
  fmt.Println(math.Sin(NUMBER))
}

