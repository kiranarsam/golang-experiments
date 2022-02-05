package main

import "fmt"

// scenario 4
func f() {
  // fmt.Println(x) // gives undefined error
}

// sceanrio 5
func f2(x int) {
  fmt.Println(x)
}

// scenario 6
var y int = 10
func f3() {
  fmt.Println(y)
}

// scenario 7
func f4() int {
  return f5()
}

func f5() int {
  return 1
}

// scenario 8
func f6() (r int) {
   r = 1
   return
}

func main() {

  xs := []float64 { 98, 93, 77, 82, 83 }
  total := 0.0
  
  for _, value := range xs {
    total += value
  }
  
  fmt.Println(total / float64(len(xs)))
  
  // scenario 2
  fmt.Println(average(xs))
  
  // scenario 3
  otherName := []float64 { 98, 93, 77, 82, 83 }
  fmt.Println(average(otherName))
  
  // scenario 4
  x := 5
  f()
  
  // scenario 5
  f2(x)
  
  // scenario 6
  f3()
  
  // scenario 7
  fmt.Println(f4())
  
  // scenario 8
  fmt.Println(f6())
  
  
}

func average(xs []float64) float64 {
  // panic("Not Implemented") // run-time error (built-in function)
  total := 0.0
  for _, value := range xs {
    total += value
  }
  
  return total / float64(len(xs))
}

