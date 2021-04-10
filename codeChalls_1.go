package main

import (
	"sort"
	"strconv"
	"strings"
)

// from codeWars (8kyu)
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

// 8kyu
func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	// your code
	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
}

// 8kyu

func Xor(a, b bool) bool {

	if a || b {
		if a && b {
			return false
		}
		return true
	}

	return false
}

// 8kyu
func CorrectTail(body string, tail rune) bool {
	return !(rune(body[len(body)-1]) != tail)
}

// 8kyu (forfeitted)
func ExpressionMatter(a int, b int, c int) int {
	expressions := []int{a + (b * c), b * a * c, c + c*b, (a + b) * c}
	sort.Ints(expressions)
	return expressions[len(expressions)-1]
}

func DNAtoRNA(dna string) string {
	return strings.Replace(dna, "T", "U", -1)
}

// from codewars 8kyu
func Summation(n int) int {
	var acum int
	for i := 0; i <= n; i++ {
		acum += i
	}
	return acum
}

// from codewars 8kyu
func Move(position int, roll int) int {
	// your code here
	return position + roll*2
}

// from codewars 8kyu
func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

// invert string, from codeWars (8kyu)
func Solution(word string) string {
	// this is how you access string positions
	// this was an 8kyu kata in codewars.com
	var (
		answer  string
		counter int = len(word) - 1
	)
	for i := 0; i < len(word); i++ {

		answer = answer + string(word[counter])
		counter = counter - 1
	}

	return answer
}

// from codewars (8kyu)
func OtherAngle(a int, b int) int {

	return 180 - (a + b)
}

//from codewars(7kyu)
func TwoToOne(s1 string, s2 string) string {
	var (
		noDuplicates string = ""
	)
	for i := 0; i < len(s1); i++ {
		if s1[i] == noDuplicates[i] {
			noDuplicates += string(s1[i])
		}
	}
	return s1 + s2
}

//from codewars (7kyu)
func NbYear(p0 int, percent float64, aug, p int) int {
	var (
		totalPop int = p0
		years    int = 0
	)
	for ; totalPop < p; years++ {
		totalPop += int((float64(totalPop)/100)*percent + float64(aug))
	}
	return years
}

// from codeWars (7kyu)
func Arithmetic(a, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	default:
		panic("Unknown operator")
	}
}

//from codewars 8kyu

func PositiveSum(numbers []int) int {
	var acum int
	for i := 0; i <= len(numbers)-1; i++ {
		if numbers[i] >= 0 {
			acum += numbers[i]
		}
	}
	return acum
}

// 8kyu

func grow(arr []int) int {
	var acum int = arr[0]
	for i := 1; i < len(arr); i++ {
		acum = acum * arr[i]
	}
	return acum
}

// 8kyu

func solve(str string) string {

	var countLowers, countUppers int

	for i := 0; i < len(str); i++ {
		if string(rune(str[i])) == strings.ToLower(string(rune(str[i]))) {
			countLowers++
		}
	}

	countUppers = len(str) - countLowers

	if countLowers > countUppers {
		return strings.ToLower(str)
	} else if countUppers > countLowers {
		return strings.ToUpper(str)
	}
	return strings.ToLower(str)
}

// seconds to readable time, from codeWars (5kyu)

func countSeconds(userInput int) string {
	// use int(f+number.decimal) to round up
	if userInput > 359999 {
		userInput = 359999
	}
	if userInput < 0 {
		userInput = -(userInput)
	}
	var (
		hours   string = strconv.Itoa(int(userInput / 3600))
		minutes string = strconv.Itoa(int(int(userInput%3600) / 60))
		seconds string = strconv.Itoa(int(int(userInput%3600) % 60))
	)

	if len(hours) == 1 {
		hours = "0" + hours
	}
	if len(minutes) == 1 {
		minutes = "0" + minutes
	}
	if len(seconds) == 1 {
		seconds = "0" + seconds
	}

	return hours + ":" + minutes + ":" + seconds
}
