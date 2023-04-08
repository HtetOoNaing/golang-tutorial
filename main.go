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

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, len(rangeOne))
	fmt.Println(rangeTwo, len(rangeTwo))
	fmt.Println(rangeThree, len(rangeThree))
	rangeOne = append(rangeOne, "IEE")
	fmt.Println(rangeOne, len(rangeOne))
}
