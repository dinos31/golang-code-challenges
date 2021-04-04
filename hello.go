package main

import (
	"fmt"
)

//remember to add formated strings with %v and %t
// remember to add for/ while loops and block declaration
//and := this weird ass syntax
// also add the scanln for input and the
// string position accessor

func main() {
	//fmt.Println("enter # of seconds")
	// var seconds int
	// fmt.Scanln(&seconds)
	// fmt.Printf("clock : %v\n", countSeconds(seconds))
	fmt.Println("enter a word")
	var myWord string
	fmt.Scanln(&myWord)
	fmt.Println(Solution(myWord))
}
