package main

import "fmt"

// more ways of declaring variables
//var i, j uint64 = 1, 2
//var c, python, java = true, false, "no!"

func main() {

	var cantidadValores int

	fmt.Println("cuantos valores desea ingresar?")
	fmt.Scanln(&cantidadValores)

	var negativos, ceros, positivos, input int

	for i := 0; i < int(cantidadValores); i++ {
		fmt.Printf("\ningrese valor # %v :", i+1)
		fmt.Scanln(&input)
		if input > 0 {
			positivos++
		}
		if input == 0 {
			ceros++
		}
		if input < 0 {
			negativos++
		}
	}

	fmt.Printf(
		"la cantidad de #'s positivos es de : %v \n\n la cantidad de #'s Negativos es de : %v \n \n la cantidad de ceros es de : %v \n ",
		positivos, negativos, ceros)
}
