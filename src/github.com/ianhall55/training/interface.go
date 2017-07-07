package main

import "fmt"

type Square struct {
  side float64
}

type Circle struct {
  radius float64
}

func (z Square) area() float64 {
  return z.side * z.side
}

func (z Circle) area() float64 {
  // pi := 3.14
  return 3.14 * z.radius * z.radius
}

type Shape interface {
  area() float64
}

func area(z Shape) {
  fmt.Println(z.area())
}

func main() {
  s1 := Square{4}
  c1 := Circle{2}
  fmt.Println(s1)
  fmt.Println(c1)
  area(s1)
  area(c1)
}
