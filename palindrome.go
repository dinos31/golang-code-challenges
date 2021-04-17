package main

// the palindrome function is case sensitive

func removeNonAlpha(str string) string {
	rts := []rune{}

	for _, value := range str {
		if value > 64 && value < 91 || value > 96 && value < 123 {
			rts = append(rts, value)
		}
	}

	return string(rts)
}

func turnAround(str string) string {
	rts := []rune{}

	for i := len(str) - 1; i >= 0; i-- {
		rts = append(rts, rune(str[i]))
	}

	return string(rts)
}

func palindrome(str string) bool {
	str = removeNonAlpha(str)
	mirror := turnAround(removeNonAlpha(str))
	return mirror == str
}
