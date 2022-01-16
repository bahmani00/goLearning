package iteration

func Repeat(chr string, repeat int) string {
	var result string
	for i := 0; i < repeat; i++ {
		result = result + chr
	}
	return result
}
