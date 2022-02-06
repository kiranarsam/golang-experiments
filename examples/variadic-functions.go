package main

import "fmt"

func sum(nums ...int) {

  fmt.Print(nums, "  ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {

  sum(10, 20)
  sum(10, 20, 30)
  sum(10, 20, 30, 40)
  
  nums := []int{1,2,3,4,5,6,7,8,9}
  sum(nums...)
}

