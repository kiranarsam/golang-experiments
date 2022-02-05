package main
import (
  "os"
  "fmt"
)

func main() {
  file, err := os.Create("test_w.txt")
  if err != nil {
    fmt.Println("Failure to open a file")
    return
  }
  defer file.Close()
  
  file.WriteString("test")
  /* ... */
  file.WriteString("\nEOF\n")
}

