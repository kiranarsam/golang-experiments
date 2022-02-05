package main 

import ("fmt"; "math")

// defining structure
type Circle struct {
  x float64
  y float64
  r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

func circleArea(c Circle) float64 {
  return math.Pi * c.r * c.r
}

// using pointer in structs type
func circleArea1(c *Circle) float64 {
  return math.Pi * c.r * c.r
}

// sceanrio 3
func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

// doing same for Rectangle
func (rec *Rectangle) area() float64 {
  return math.Sqrt(rec.y2 - rec.y1) + math.Sqrt(rec.x2 - rec.x1) 
}

// main function
func main() {
  var c Circle
  r := new(Rectangle)
  
  c1 := Circle { x: 0, y: 0, r: 5 }
  c2 := Circle{ 0, 0, 5 }
  
  fmt.Println(c.x, c.y, c.r)
  fmt.Println(r.x1, r.y1, r.x2, r.y2)
  fmt.Println(c1.x, c1.y, c1.r)
  fmt.Println(c2.x, c2.y, c2.r)
  
  c2.x = 10
  c2.y = 10
  c2.r = 15
  fmt.Println(c2.x, c2.y, c2.r)
  
  fmt.Println(circleArea(c2)) // sceanrio 1
  fmt.Println(circleArea1(&c2)) // sceanrio 2 with pointer and strcts type

  // sceanrio 3, using simplest method
  fmt.Println(c2.area())
  
  r.x1, r.y1, r.x2, r.y2 = 0, 0, 10, 10
  fmt.Println(r.area())
  
  r1 := Rectangle{0, 0, 20, 20}
  fmt.Println(r1.area())
  
}

