package main

import "fmt"

func main() {

  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)
  
  output := input * 2
  
  fmt.Println(output)
  
  x := 5; x += 1
  fmt.Println(x)
  
  // convert Fahrenheit into Celsius
  
  fmt.Print("Enter a temparature in Fahrenheit: ")
  var temp float64
  fmt.Scanf("%f", &temp)
  
  temp = (temp - 32) * (5.0 / 9)
  fmt.Println("Temp in Celsius = ", temp)
  
  const ONE_FEET float64 = 0.3048
  fmt.Println("2 feets = ", (2 * ONE_FEET))
  

}
