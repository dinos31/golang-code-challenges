package main

import (
	"strconv"
)

func Solution(word string) string {
	// this is how you access string positions
	// this was an 8kyu kata in codewars.com
	var (
		answer  string
		counter int = len(word) - 1
	)
	for i := 0; i < len(word); i++ {

		answer = answer + string(word[counter])
		counter = counter - 1
	}

	return answer
}
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

	if len(hours) == 1 {
		hours = "0" + hours
	}
	if len(minutes) == 1 {
		minutes = "0" + minutes
	}
	if len(seconds) == 1 {
		seconds = "0" + seconds
	}

	return hours + ":" + minutes + ":" + seconds
}
