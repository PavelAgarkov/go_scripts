package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AZC Celsius = -273.15
	FC  Celsius = 0
	BC  Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}
