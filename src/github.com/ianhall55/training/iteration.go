package main

import "fmt"

func main() {
  i := 0
  for i < 1000 {
    fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
    i++
  }
}
