package main

import (
	"fmt"

	"rsc.io/quote"
)

// initialized variable with type declaration
// var goint int = 3
// var gostring string = "this string is strongly typed"
// var gofloat float32 = 23

//there are two ways of declaring multiple variables at the time
// var (
// 	myString string = "this is my string"
// 	myInt    int    = 123
// ) // var block

/* you can also declare variables like this*/

func main() {
	// this syntax detects variable type, but it has
	// to be written and USED inside a function
	//weirdSyntax := `some weird syntax`
	// fmt.Println("\n" + quote.Opt() + "\n") // this is straing F.A.C.T.S
	// fmt.Println(gostring + "\n")
	// fmt.Println(weirdSyntax)
	// the following is a formated print
	// you can also use a %T to print the variable's type
	fmt.Printf("\n%v\n", quote.Opt())
	// inside a formated string
	//fmt.Printf("%v , %T \n", gofloat, gofloat)
	//fmt.Printf("%v,%v", myString, myInt)
	fmt.Println(countSeconds(-12340))
}
