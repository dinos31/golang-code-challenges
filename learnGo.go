package main

import (
	"fmt"
)

// more ways of declaring variables
//var i, j uint64 = 1, 2
//var c, python, java = true, false, "no!"

func main() {
	fmt.Printf("sin especiales : %v\n", removeNonAlpha("Mateo"))
	fmt.Printf("al contrario : %v\n", removeNonAlpha(turnAround("Mateo")))
	fmt.Println(palindrome("Mateo"))
}
