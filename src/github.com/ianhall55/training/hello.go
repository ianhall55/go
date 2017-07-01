package main

import "fmt"

func main() {
  x := 'b'

  y := 5
  a := 12
  
  fmt.Printf("%v \n", x)
  fmt.Printf("%v \n", y)

   z := foo(y, a)
   fmt.Println(z)
}

func foo(x, y int) int {
  fmt.Println(x + y)
  return 7
}
