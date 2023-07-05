// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

import "fmt"

type Celsius float64      // 摄氏温度
type Fahrenheit float64   // 华氏温度

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC Celsius = 0
  BoilingC Celsius = 100
)

func CtoF (c Celsius) Fahrenheit {return Fahrenheit (c*9/5 + 32)}
func FtoC (f Fahrenheit) Celsius {return Celsius ((f-32)*5/9)}

func main () {
  c := FtoC(212.0)
  fmt.Println(c)
}