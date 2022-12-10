package main

import (
	"fmt"
	"headfirstgo/src/chapter4/calc"
	"headfirstgo/src/chapter4/greeting"
)

func main() {
	//chapter2.PassFail()
	//chapter2.FileSize()
	//chapter2.Guess()
	//chapter3.Printf()
	//chapter3.DisplayArea()
	//chapter3.RepeatLine("hello", 3)
	//chapter4.Tocelsius()
	greeting.Greeting()
	greeting.Hello()
	greeting.Hi()
	fmt.Println("1 + 2 = ", calc.Add(1, 2))
	fmt.Println("1 - 2 = ", calc.Dec(1, 2))
}
