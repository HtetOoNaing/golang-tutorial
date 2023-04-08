package main

import "fmt"

func main() {
	var ages [3]int = [3]int{10, 20, 30}
	names := [4]string{"HEE", "JEE", "MEE", "KEE"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 58)

	fmt.Println(scores, len(scores))
}
