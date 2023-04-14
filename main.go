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

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fmt.Println(name, price)
	case "s":
		fmt.Println("You chose s")
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	default:
		fmt.Println("That is not a valid option...")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
	// fmt.Println(myBill)
	// fmt.Println(myBill.format())
}
