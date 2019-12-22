// Package hamming is a package to implement the calculation of dna
package hamming

import "errors"

// Distance calculates the hamming distance
func Distance(a, b string) (dis int, err error) {
	if len(a) != len(b) {
		return 0, errors.New("Not Equal length")
	}

	for index := 0; index <= len(a)-1; index++ {
		if a[index] != b[index] {
			dis++
		}
	}

	return
}
