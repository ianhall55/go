package main

import "fmt"

func main() {
  fmt.Println(half(10))
  fmt.Println(half(15))
}

func half(x int) (float64, bool) {
  return (float64(x) / 2), (x % 2 == 0)
}
