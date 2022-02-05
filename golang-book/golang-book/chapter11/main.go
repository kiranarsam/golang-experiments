package main

import "fmt"
import "golang-book/chapter11/math"

// alias name
import m "golang-book/chapter11/math"

func main() {

  xs := []float64{1, 2, 3, 4}
  avg := math.Average(xs)
  fmt.Println(avg)

  // m is the alias
  avg1 := m.Average(xs)
  fmt.Println(avg1)
}

