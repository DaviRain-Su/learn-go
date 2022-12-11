package chapter6

import (
	"fmt"
	"math"
)

func DisplaySlice() {
	mySlice := make([]string, 7)
	mySlice[0] = "do"
	mySlice[1] = "hello, world!"
	fmt.Println("display mySlice is ", mySlice)
	fmt.Println("mySlice len is ", len(mySlice))
	for _, letter := range mySlice {
		fmt.Println("slice is ", letter)
	}
	primes := []int{1, 2, 3, 4, 5}
	for _, value := range primes {
		fmt.Println("value is ", value)
	}
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3]
	fmt.Println("slice1 is = ", slice1)
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))

	s1 := []string{"s1", "s1"}
	s1 = append(s1, "s2", "s2")
	s1 = append(s1, "s3", "s3")
	s1 = append(s1, "s4", "s4")
	fmt.Println("s1 is ", s1)

	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])

	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice)
	fmt.Println("intSlice len is ", len(intSlice))
	fmt.Println("stringSlice len is ", len(stringSlice))
	severalInts(1)
	severalInts(1, 2, 3)
	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d", "e")
	severalStrings()
	mix(1, true, "a", "b")
	mix(2, false, "a", "b", "c", "d")
	fmt.Println("max is ", maximum(1, 2, 3, 4, 5))
	fmt.Println("max is ", maximum(-1, -2, -3, -4, -5))
}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func Average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
