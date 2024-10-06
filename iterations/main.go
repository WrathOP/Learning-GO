package iteration

func Repeat(character string, repeat int) string {
	repeated := ""
	for i := 0; i <= repeat; i++ {
		repeated += character
	}

	return repeated
}
