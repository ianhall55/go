package main

import "fmt"

type person struct {
  first string
  last  string
  age   int
}

func main() {
  p1 := person{"Kasey", "Kozitza", 28}
  p2 := person{"Ian", "Hall", 27}
  fmt.Println(p1)
  fmt.Println(p2)
}
