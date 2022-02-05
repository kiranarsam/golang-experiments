package main

import "fmt"

// scenario 1  (Has-a)
type Person struct {
  Name string
}

func (p *Person) Talk() {
  fmt.Println("Hi, my name is ", p.Name)
}

// Scenario 2 (Has-a) Not to be preferred
type Android struct {
 Person Person
 Model string
}

// Scenario 3 (Is-a) Preferred one
type AndroidA1 struct {
  Person
  Model string
}

func main() {
  // sceanrio 1
  p1 := new(Person)
  p1.Name = "John"
  p1.Talk()
  
  // scenario 2
  a1 := new(Android)
  a1.Person.Name = "John-Android"
  a1.Person.Talk()
  
  // scenario 3
  a2 := new(AndroidA1)
  a2.Person.Name = "John-AndroidA1"
  a2.Person.Talk()
  // call direcltly
  a2.Talk() // Is-a relation (Preferred one)
    

}

