package main

import (
	"math"
	"strings"
)

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

// 8kyu
func HowManyDalmatians(number int) string {

	dogs := []string{"Hardly any", "More than a handful!", "Woah that's a lot of dogs!", "101 DALMATIONS!!!"}

	if number <= 10 {
		return dogs[0]
	}
	if number <= 50 {
		return dogs[1]
	}
	if number < 101 {
		return dogs[2]
	}

	return dogs[3]
}

// 8kyu
func GetSize(w, h, d int) [2]int {

	return [2]int{2*(h*w) + 2*(h*d) + 2*(w*d), w * h * d}
}

func AbbrevName(name string) string {
	firstInitial := string(name[0])
	secondInitial := ""
	for i := 0; i < len(name); i++ {
		if name[i] == ' ' {
			secondInitial = string(rune(name[i+1]))
		}
	}
	return strings.ToUpper(firstInitial + "." + secondInitial)
}

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func QuarterOf(month int) int {
	return int(math.Ceil(float64(month) / 3))
}
