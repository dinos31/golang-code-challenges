package main

import (
	"fmt"

	"rsc.io/quote"
)

// initialized variable with type declaration
var goint int = 3
var gostring string = "this string is strongly typed"

/* you can also declare variables like this*/

func main() {
	// this syntax detects variable type, but it has
	// to be written and USED inside a function
	weirdSyntax := `some weird syntax`

	// fmt.Println("\n" + quote.Opt() + "\n") // this is straing F.A.C.T.S
	// fmt.Println(gostring + "\n")
	// fmt.Println(weirdSyntax)
	// the following is a formated print
	fmt.Printf("%v\n%v used to declare a %T ", quote.Opt(), weirdSyntax, weirdSyntax)
	// you can also use a %T to print the variable's type
	// inside a formated string
}
