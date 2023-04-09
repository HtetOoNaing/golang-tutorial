package main

import (
	"fmt"
)

var score = 99.5

func main() {
	sayHello("mario")

	// get points from greetings.go (package scope)
	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}
