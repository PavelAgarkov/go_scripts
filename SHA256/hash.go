package SHA256

import (
	"crypto/sha256"
	"fmt"
)

func Hash256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	count := countDifferenceBytes(c1, c2)

	fmt.Printf("%d\n", count)
}

func countDifferenceBytes(hashOne [32]byte, hashTwo [32]byte) int {
	bitCounter := 0
	for i, v := range hashOne {
		if v != hashTwo[i] {
			bitCounter++
		}
	}
	return bitCounter
}
