package main

import "fmt"

func main() {
	// strings
	// var nameOne string = "Htet"
	// var nameTwo = "Oo"
	// var nameThree string
	// fmt.Println(nameOne, nameTwo, nameThree)
	// nameOne = "Htet1"
	// nameTwo = "Oo1"
	// nameThree = "Naing1"
	// fmt.Println(nameOne, nameTwo, nameThree)
	// nameFour := "luios"
	// fmt.Println(nameFour)

	// ints
	// var ageOne int = 20
	// var ageTwo = 21
	// ageThree := 22
	// fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	// var numOne int8 = 127
	// var numTwo int8 = -128
	// var numThree uint8 = 255
	// fmt.Println(numOne, numTwo, numThree)

	// floats
	// var scoreOne float32 = 25.43
	// var scoreTwo float64 = 3423423423423424234234234234.32432
	// scoreThree := 2.6 // float64
	// fmt.Println(scoreOne, scoreTwo, scoreThree)

	name := "Htet"
	age := 20
	// Print
	fmt.Print("hello, ")
	fmt.Print("GO \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("Hello, GO")
	fmt.Println("My name is", name, "and my age is", age)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("My name is %v and my age is %v \n", name, age)
	fmt.Printf("My name is %q and my age is %q \n", name, age)
	fmt.Printf("Age is type %T \n", age)
	fmt.Printf("you score %f points! \n", 255.55)
	fmt.Printf("you score %0.1f points! \n", 255.55)
}
