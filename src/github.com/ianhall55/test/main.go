package main

import "fmt"

type person struct {
  fname string
  lname string
}

func main() {

  xi := []int{2,3,4,5,4}
  fmt.Println(xi)

  m := map[string]int{
    "Todd": 45,
    "Ian": 27,
  }
  fmt.Println(m)

  p1 := person{
    "Ian",
    "Hall",
  }

  fmt.Println(p1)
}
