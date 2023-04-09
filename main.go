package main

import (
	"fmt"
)

var score = 99.5

func main() {
	myBill := newBill("mario's bil")
	myBill.updateTip(10.5)
	myBill.addItem("chili", 44)
	fmt.Println(myBill)
	fmt.Println(myBill.format())
}
