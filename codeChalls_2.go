package main

// 8kyu
func OddCount(n int) int {
	return n / 2
}

func IsUpperCase(s string) bool {
	x := []rune(s)
	for _, index := range x {
		if index >= 97 || 122 <= index {
			return false
		}
	}
	return true
}
