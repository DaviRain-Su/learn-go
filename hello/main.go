package main

import "fmt"
import temparray "hello/array"

// main
func main() {

	fmt.Println("Hello,world")

	exampleArray()

	temparray.Array()
}

func exampleArray() {

	var array [5]int

	fmt.Println(array)

	array_1 := [...]int{10, 20, 30, 40, 50, 0, 10, 20}

	array_2 := [5]int{10, 20, 30, 40, 50}

	array_3 := [5]int{1: 10, 2: 10}

	fmt.Println(array_1)

	fmt.Println(array_2)

	fmt.Println(array_3)

	array_3[2] = 23

	fmt.Println(array_3)

	array_4 := [5]*int{0: new(int), 1: new(int)}

	*array_4[0] = 10

	*array_4[1] = 20

	fmt.Println(array_4)
	fmt.Println(*array_4[0])

	for index, value := range array_4 {
		if value != nil {
			fmt.Println("index = ", index, "value = ", *value)
		}
	}

	var array1 [5]string

	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	array1 = array2

	fmt.Println("array1 = ", array1)

	fmt.Println("array2 = ", array2)

	var array3 [5]string

	array4 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	array3 = array4

	fmt.Println(array3)

	var array5 [3]*string

	array6 := [3]*string{new(string), new(string), new(string)}

	*array6[0] = "Red"
	*array6[1] = "Blue"
	*array6[2] = "Green"

	array5 = array6

	fmt.Println(array5)

	for _, value := range array6 {
		fmt.Println("array6 = ", *value)
	}

	for _, value := range array5 {
		fmt.Println("array5 = ", *value)
	}

}
