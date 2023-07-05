// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC Celsius = 0
  BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%gC", c)}

func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f)}

func (K Kelvin) String() string { return fmt.Sprintf("%gK", K)}

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5+32)}

func FtoC(f Fahrenheit) Celsius { return Celsius((f-32)*5/9)}

func CtoK(c Celsius) Kelvin { return Kelvin(c+273.15)}

func KtoC(K Kelvin) Celsius { return Celsius(K-273.15)}
