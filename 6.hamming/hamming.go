package hamming

import (
	"fmt"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		err := "You got different length"
		return 0, fmt.Errorf(err)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
