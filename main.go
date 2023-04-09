package main

import (
	"fmt"
)

var score = 99.5

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	// looping maps
	for k, v := range menu {
		fmt.Println("Key is", k, "and value is", v)
	}

	phonebook := map[int]string{
		12343243: "meria",
		53453433: "luios",
		53452322: "martin",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[53453433])
	phonebook[53453433] = "bowser"
	fmt.Println(phonebook)
}
