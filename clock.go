package main

import (
	"fmt"
	"strconv"
)

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

	fmt.Printf("\n%v hour, %v minutes and %v seconds\n",
		hours, minutes, seconds)

	return ""
}
