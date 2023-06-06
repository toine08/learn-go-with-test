package main

import (
	"fmt"
)

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	var times int
	var char string

	fmt.Print("What characters do you want?")
	fmt.Scanf("%s", &char)
	fmt.Print("How many times do you want it?")
	fmt.Scanf("%d", &times)

	repeated := Repeat(char, times)
	fmt.Printf("You chose to repeat '%s' %v times: %s\n", char, times, repeated)
}
