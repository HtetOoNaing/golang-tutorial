package main

import "fmt"

func main() {
	var ages [3]int = [3]int{10, 20, 30}
	names := [4]string{"HEE", "JEE", "MEE", "KEE"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))
}
