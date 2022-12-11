package chapter6

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Average2() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("Arverage : %0.2f\n", Average(numbers...))
}
