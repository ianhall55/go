package main

import "fmt"

func main() {
  mySlice := []int{1,2,3,4,5,6,7}
  for _, el := range mySlice {
    if (el % 2 == 1) {
      fmt.Println(el)
    }
  }
  fmt.Println(mySlice[2:])
  fmt.Println(len(mySlice))
  fmt.Println(cap(mySlice))

  slice2 := make([]string, 5)
  fmt.Println(slice2)
  fmt.Println(len(slice2))
  fmt.Println(cap(slice2))
  slice2[0] = "Pip"
  slice2 = append(slice2, "ian")
  fmt.Println(slice2)
}


// slices

// initialization
// shorthand
// :=
// var
// var
// make(type, length, underlying arr capacity)
// make([]int, 0, 5)
// multi-initial array of array of ints
// make([][]int)
