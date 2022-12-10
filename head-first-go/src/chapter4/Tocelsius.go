package chapter4

import (
	"fmt"
	"headfirstgo/src/keyboard"
	"log"
)

func Tocelsius() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degress Celsius\n", celsius)
}
