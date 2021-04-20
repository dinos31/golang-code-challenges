package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func Operaciones() string {

	var (
		operaciones int
		suma        int
	)

	fmt.Println("cuantas operaciones desea realizar?")
	fmt.Scanln(&operaciones)
	for i := 0; i < operaciones; i++ {
		fmt.Printf("ingrese valor del numero %v \n", i+1)
		var num int
		fmt.Scanln(&num)
		suma += num
	}
	fmt.Printf("el total es %v", suma)

	return ""
}

func calcularSalario(salario, aumento float64) float64 {
	return float64(salario + (salario / 100 * aumento))
}

func promedio(notas []float32) string {
	var suma float32 = 0
	for index := 0; index < len(notas); index++ {
		suma += notas[index]
	}

	var avg float32 = suma / float32(len(notas))
	if avg >= 3.0 {
		return fmt.Sprintf("aprobo con : %f", avg)
	}
	return fmt.Sprintf("reprobo con : %f", avg)
}

func countSheep(num int) string {
	var sheeps string = ""
	var counter int = 1
	for counter <= num {
		sheeps += strconv.Itoa(counter) + " sheep..."
		counter++
	}
	return sheeps
}

func GetPlanetName(ID int) string {

	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"
	default:
		return "Pluto" // ;-)
	}
}

func monkeyCount(n int) []int {
	var countedMonkeys = []int{}
	for len(countedMonkeys) != n {
		countedMonkeys = append(countedMonkeys, len(countedMonkeys)+1)
	}
	return countedMonkeys
}

func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}
func century(year int) int {

	return int(math.Ceil(float64(year) / 100))
}
func Past(h, m, s int) int {
	return (h * 3600000) + (m * 60000) + (s * 1000)
}

func MakeUpperCase(str string) string {
	runes := []rune{}

	for _, value := range str {
		runes = append(runes, unicode.ToUpper(value))
	}
	return string(runes)
}

func SquareSum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += int(math.Pow(float64(value), 2))
	}
	return sum
}

func SubtractSum(n int) string {
	fruits := map[int]string{}

	x := strings.Split(strconv.Itoa(n), "")
	y := make([]int, len(x))
	acum := 0

	for _, value := range x {
		newNumber, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		} else {
			y = append(y, newNumber)
		}
	}
	for _, value := range y {
		acum += value
	}

	return fruits[acum]
}
