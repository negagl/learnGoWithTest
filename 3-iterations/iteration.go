package iterations

func Repeat(character string, times int) (repeatedString string) {
	for i := 0; i < times; i++ {
		repeatedString += character
	}

	return
}