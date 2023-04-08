package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "hello there my friends!"
	fmt.Println(strings.Contains(greeting, "helloo"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))
}
