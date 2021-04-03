package main

import (
	"fmt"

	"rsc.io/quote"
)

//initialized variable with type declaration
var goint int
var gostring string = "this string is strongly typed"
var gofloat float32 = 23

//there are two ways of declaring multiple variables at the time
var (
	myString string = "string inside block declarations"
	myInt    int    = 123
) //var block

// this is the equivalent of a while loop
// for goint < 10 {
// 	fmt.Println(goint + 1)
// 	goint++
// }
// classic for loop in go
//for i := 0; i < 10; i++{
//   your code here
//}

func main() {
	// this syntax detects variable type, but it has
	// to be written and USED inside a function
	//weirdSyntax := `some weird syntax`

	// the following is a formated print
	// formated text keyLetters v = value, T = variable's type
	// you can also use a %T to print the variable's type
	fmt.Printf("\n%v\n",
		quote.Opt())

	fmt.Println("enter # of seconds")
	var seconds int
	fmt.Scanln(&seconds)

	fmt.Printf("clock : %v", countSeconds(seconds))
}
