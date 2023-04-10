package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var score = 99.5

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func main() {
	myBill := createBill()
	fmt.Println(myBill)
	fmt.Println(myBill.format())
}
