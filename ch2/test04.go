package main

import(
  "fmt"
  "os"
  "strconv"
)

func gcd(x,y int) int {
    for y != 0 {
      x, y = y, x%y
    }
  return x
  }
  
func main () {
  if len(os.Args) < 3 {
    fmt.Println("请输入两个整数")
    return
  }

  x, err := strconv.Atoi(os.Args[1])
  if err != nil {
    fmt.Println("第一个参数无效：", err)
    return
  }

  y, err := strconv.Atoi(os.Args[2])
  if err != nil {
    fmt.Println("第二个参数无效：", err)
    return
  }

  result := gcd(x, y)
  fmt.Println("最大公约数：", result)
}