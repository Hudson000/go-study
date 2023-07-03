package main

import "fmt"

func main() {
  func incr(p*int) int {
    *p++
    return *p
    }

v := 1
incr(&v)
fmt.Println(incr(&v))
  }