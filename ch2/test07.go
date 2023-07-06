package main

import (
  "fmt"
  "test/popcount"
  "os"
  "strconv"
)

func main() {
  x := os.Args[1]
  y, err := strconv.ParseUint(x,10,64)
  if err != nil {
    return
  } 
  fmt.Println(popcount.PopCount(y))
}