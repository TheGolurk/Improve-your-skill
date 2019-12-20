// Package hamming is a package to implement the calculation of dna
package hamming

// Distance calculates the hamming distance
func Distance(a, b string) (dis int, err error) {
	if len(a) != len(b) {
		return 0, nil
	}

	dna := make(map[int]byte)

	for index := 0; index <= len(a)+len(b); {
		dna[index] = a[index]
	}
	
	return
}
