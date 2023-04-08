package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	// sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 900)
	fmt.Println(index)
}
