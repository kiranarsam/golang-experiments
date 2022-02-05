package main

import "fmt"

// scenaio 1
func zero(x int) {
  x = 0
}

// scenario 2 (using pointers)
func update(xPtr *int) {
  *xPtr = 0
}

// scenario 3 (using "new" built-in function)
func one(xPtr *int) {
  *xPtr = 1
}

// scenario 4
func square(x *float64) {
  *x = *x * *x
}

// sceanrio 5
func swap(x *int, y *int) {
   temp := *x
   *x = *y
   *y = temp
}


func main() {
  // scenario 1
  x := 5
  zero(x)
  fmt.Println(x) // still x is 5
  
  // scenario 2
  update(&x)
  fmt.Println(x) // Now x is 0

  // scenario 3 Using "new"
  xPtr := new(int)
  one(xPtr)
  fmt.Println(*xPtr) // x is 1  

  // scenario 4
  a := 1.5
  square(&a)
  fmt.Println(a)
  
  // scenario 5
  l := 1; m := 2
  swap(&l, &m)
  fmt.Println(l, m)
}


