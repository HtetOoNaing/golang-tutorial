package main

import (
	"fmt"
	"strings"
)

func circleArea(s string) (string, string) {
	upper := strings.ToUpper(s)
	names := strings.Split(upper, " ")
	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"

}

func main() {
	fn1, sn1 := circleArea("dj soda")
	fmt.Println(fn1, sn1)
}
