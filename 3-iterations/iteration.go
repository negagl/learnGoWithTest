package iterations

func Repeat(character string, times int) string {
	var repeatedString string

	for i := 0; i < times; i++ {
		repeatedString = repeatedString + character
	}

	return repeatedString
}