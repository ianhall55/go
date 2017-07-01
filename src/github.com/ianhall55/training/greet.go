package main

import "fmt"

func main() {
  data := []float64{12, 14, 15, 17, 42}
  fmt.Println(average(data...))
}

func greet(fname, lname string) string {
  return fmt.Sprint(fname, lname)
}

func average(sf ...float64) float64 {
  fmt.Println(sf)
  var total float64
  for _, v := range sf {
    total += v
  }
  fmt.Println(total)
  return total / float64(len(sf))
}
