package main

// from codeWars (8kyu)
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

// from codewars (8kyu)
func OtherAngle(a int, b int) int {

	return 180 - (a + b)
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
