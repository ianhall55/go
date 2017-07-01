package main

import "fmt"

var x int

func main() {
  fmt.Println(increment())
  fmt.Println(increment())
  fmt.Println(x)
}

func increment() int {
  x++
  return x
}
