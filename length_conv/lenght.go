package length_conv

import (
	"fmt"
	"os"
	"strconv"
)

type Meter float64
type Foot float64

const (
	METER_TO_FOOT = 3.28084
)

func (m Meter) MeterToFoot() Foot {
	return Foot(m * METER_TO_FOOT)
}

func (f Foot) FootToMeter() Meter {
	return Meter(f / METER_TO_FOOT)
}

func StartConvert() {
	for _, argErr := range os.Args[1:] {
		t, err := strconv.ParseFloat(argErr, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		m := Meter(t).MeterToFoot()
		f := Foot(t).FootToMeter()

		fmt.Printf("%v foots = %v meters , %v foots = %v meters \n", m, t, t, f)
	}
}
