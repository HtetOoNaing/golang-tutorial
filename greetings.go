package main

import "fmt"

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("Hello", n)
}

func showScore() {
	//get score from main.go (packeage scope)
	fmt.Println("Your score is", score)
}
