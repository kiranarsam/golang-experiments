package main

import (
        "fmt"
//        "math"
       )

// interface
type Shape interface {
  area() float64
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  
  return area
}

type Circle struct {
  x1, y1, r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
  return c.r
}

func (r *Rectangle) area() float64 {
  return r.x2 + r.y2
}

// scenario 2
type MultiShape struct {
  shapes []Shape
}

func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}


func main() {

  c := Circle{0, 0, 10.0}
  r := Rectangle{0, 0, 30.0, 30.0}

  fmt.Println(totalArea(&c, &r))

  c1 := Circle{0, 0, 20.0}
  r1 := Rectangle{0, 0, 60.0, 60.0}  
  var m MultiShape
  m.shapes = make([]Shape, 2)
  m.shapes[0] = &c1
  m.shapes[1] = &r1
  fmt.Println(m.area())
}



