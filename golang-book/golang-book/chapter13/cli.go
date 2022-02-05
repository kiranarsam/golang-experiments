package main

import ("fmt"; "flag"; "math/rand")

func main() {
  // Define flags
  maxp := flag.Int("max", 6, "the max value")
  
  // Pasrse
  flag.Parse()
  
  // Generate a number between 0 and max
  fmt.Println(int(*maxp))
  fmt.Println(rand.Intn(*maxp))
}

