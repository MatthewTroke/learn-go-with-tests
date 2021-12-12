package iteration

// Repeat takes a character and concats it x times
func Repeat(character string, x int) (repeated string) {
	for i := 0; i < x; i++ {
		repeated += character
	}

	return
}
