package main

import "fmt"

func main() {

  // scenario 1
  add := func(x, y int) int {
           return x + y
         }
  
  fmt.Println(add(1,1))
  
  // scenario 2 (Here, x and increment function becomes closure
  x := 0
  increment := func() int {
    x++
    return x
  }
  fmt.Println(increment())
  fmt.Println(increment())
  
  // scenario 3
  nextEven := makeEvenGenerator()
  fmt.Println(nextEven())  // 0
  fmt.Println(nextEven())  // 2
  fmt.Println(nextEven())  // 4

  //sceanrio 4 (Recursion)
  fmt.Println(factorial(5))  
}

// scenario 3 (closure function, i and return func, becomes closure)
func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

// scenario 4 (Recursion)
func factorial( x uint ) uint {
  if x == 0 {
    return 1
  }
  
  return x * factorial(x-1)
}

