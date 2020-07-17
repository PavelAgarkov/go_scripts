package bytesBuffer

import (
	"bytes"
	"fmt"
)

func Initialization(arr []string) {
	newStr := comma(arr)
	fmt.Print(newStr)
}

func comma(str []string) string {
	var buf bytes.Buffer
	counter := 0

	for _, v := range str {

		if counter%3 == 0 && counter > 0 {
			buf.WriteString(", ")
		}

		buf.WriteString(v + " ")
		counter++
	}
	return buf.String()
}
