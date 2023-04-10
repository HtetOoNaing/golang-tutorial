package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var score = 99.5

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Create a new bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func main() {
	myBill := createBill()
	fmt.Println(myBill)
	fmt.Println(myBill.format())
}
