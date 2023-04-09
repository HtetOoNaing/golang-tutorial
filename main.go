package main

import (
	"fmt"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("x value is", x)
	// 	x++
	// }

	// for x := 0; x < 5; x++ {
	// 	fmt.Println("x value is", x)
	// }

	names := []string{"mario", "luigi", "yoshi", "peach"}
	// for x := 0; x < len(names); x++ {
	// 	fmt.Println(x, "is", names[x])
	// }
	// for index, value := range names {
	// 	fmt.Printf("The value at index %v is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}
}
