package main

import "fmt"

func main() {
  //var x  [5]int
  var x [5]float64
  x[4] = 100
  fmt.Println(x)
  fmt.Println(x[4])
  x[0] = 10
  x[1] = 20
  x[2] = 32
  x[3] = 40
  
  var total float64 = 0
  for i := 0; i < 5; i++ {
    total += x[i]
  }
  fmt.Println(total/float64(len(x)))

  total = 0
  for _, value := range x {  // _ A single underscore is used to tell the compiler that we don't need iterator variable
    total += value
  }
  fmt.Println(total/float64(len(x)))
  
  // slices
  
  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)
  
  slice1 = []int{1,2,3}
  slice2 = make([]int, 2)
  copy(slice2, slice1)
  fmt.Println(slice1, slice2)
  
  slice1 = make([]int, 5, 10)
  fmt.Println("Len of slice = ", len(slice1))
  slice2 = append(slice1, 1,2,3,4)
  copy(slice1, slice2)
  fmt.Println("Len of slice = ", len(slice1))
  fmt.Println("Len of slice = ", slice1)
  fmt.Println("Len of slice = ", slice2)

  // maps ** we should always initialize **
  
  var mapDict map[string]int = make(map[string]int)
  mapDict["key"] = 10
  fmt.Println(mapDict)
  mapDict["key2"] = 100
  mapDict["key3"] = 200
  fmt.Println(len(mapDict))
  fmt.Println(mapDict["key2"])
  delete(mapDict, "key2")
  fmt.Println(mapDict)
  
  name, ok := mapDict["key10"]
  fmt.Println(name, ok)
  
  if name, ok := mapDict["key3"]; ok {
    fmt.Println(name, ok)
  }
  
  elements := map[string]map[string]string {
    "H" : map[string]string {
            "name" : "Hydrogen",
            "state" : "gas",
     },
     "He" : map[string]string {
             "name" : "Helium",
             "state" : "gas",
     },
   }
   
   if el, ok := elements["He"]; ok {
     fmt.Println(el["name"], el["state"])
   }
  
  
    
  

}

