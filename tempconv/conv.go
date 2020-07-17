package tempconv

import (
	"fmt"
	"os"
	"strconv"
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func StartConvert() {
	for _, argErr := range os.Args[1:] {
		t, err := strconv.ParseFloat(argErr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))
	}
}
