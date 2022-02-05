package main

import "fmt"

// scenario 1
func f() (int, int) {
  return 5, 6
}

// scenario 2
func add(args ...int) int {
  total := 0
  for _, value := range args {
    total += value
  }
  return total
}
/*
prototype for Println
func Println(a ...interface{}) (n int, err error)
*/

func main() {
  // scenario 1
  x, y := f()
  fmt.Println(x,y)  
  
  // scenario 2
  fmt.Println(add(1,2,3,4,5))
  
  // scenario 3 (slice)
  xs := []int{1,2,3,4,5,6}
  fmt.Println(add(xs...))
  
}

