package main

import "fmt"

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
