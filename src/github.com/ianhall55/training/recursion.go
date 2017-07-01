package main

import "fmt"

func main() {
  fmt.Println(factorial(5))
  fmt.Println(factorial(7))
}

func factorial(x int) int {
  if x == 1 {
    return x
  }
  return  x * factorial(x - 1)
}
