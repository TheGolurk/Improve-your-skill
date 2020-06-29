package reverse

func Reverse(word string) string {
	cpi := word
	for i, val := range word {
		cpi[len(word)-i] = byte(val)
	}
	return cpi
}
