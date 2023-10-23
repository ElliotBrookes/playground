package iteration

func Repeat(value string, count int) string {
	var result string

	for i := 0; i < count; i++ {
		result += value
	}

	return result
}
