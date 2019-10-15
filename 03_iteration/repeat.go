package iteration

// Repeat returns string with char repeated repeatCount times
func Repeat(char string, repeatCount int) string {
	var result string
	for i := 0; i < repeatCount; i++ {
		result += char
	}
	return result
}
