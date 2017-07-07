package main

import "fmt"

type Car struct {
  make    string
  model   string
  noise   string
  honkStr int
}

type SportsCar struct {
  Car
}

func (car Car) honk() string {
  honkString :=  ""
  i := 0
  for i < car.honkStr {
    honkString += car.noise + " "
    i++
  }
  return honkString
}

func (s SportsCar) honk() string {
  return "VROOM!"
}

func main() {
  c1 := Car{"Honda", "Civic", "Beep", 5}
  s1 := SportsCar{Car: Car{"Audi", "A5", "Beep", 2}}
  fmt.Println(c1.honk())
  fmt.Println(s1.honk())
  fmt.Println(s1.Car.honk())
}
