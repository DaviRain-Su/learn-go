package chapter1

import (
	"fmt"
	"math"
	"strings"
)

func DefineValue() {
	fmt.Println("Hello, World!")
	fmt.Println("First argument", "Second argument")
	fmt.Println("Another line")
	var num = math.Floor(2.75)
	fmt.Println("num is ", num)
	var tempStr = strings.Title("hello, Rust!")
	fmt.Println("temp_str is ", tempStr)
	fmt.Println("Hello, \nGo!")
	fmt.Println("Hello, \tGo!")
	fmt.Println("Hello, \"\"")
	fmt.Println("Hello, \\")

	// define variable
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"

	// use variable
	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	var myInt int
	var myFloat float64
	var myString string
	var myBool bool
	fmt.Println(myInt, myFloat, myString, myBool)
}
