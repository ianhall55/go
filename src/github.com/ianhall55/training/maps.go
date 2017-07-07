package main

import "fmt"

func main() {
  m := make(map[string]int)
  m["a"] = 1
  m["c"] = 2
  fmt.Println(m)

  n := map[string]int{"foo": 1, "bar": 3}
  fmt.Println(n)

  if val, exists := n["foo"]; exists {
    fmt.Println(val)
    fmt.Println(exists)
  }

  for k, v := range n {
    fmt.Println(k, ":", v)
  }

  i, exists := m["a"]
  if exists {
    fmt.Println(i)
  }
  i, exists2 := m["z"]
  if !exists2 {
    fmt.Println(i)
  }
}
