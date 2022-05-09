package goinaction

import "fmt"

/// - 数组的内部实现和基础功能
/// - 使用切片管理数据集合
/// - 使用映射管理键值对
func Array() {
	fmt.Println("Test Array")
}

func DeclineArray() {
	var array [5]int
	fmt.Println("array = ", array)

	array1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println("array1 = ", array1)

	array2 := [...]int{10, 20, 30, 40}
	fmt.Println("array2 = ", array2)

	array3 := [5]int{1: 10, 2: 20}
	fmt.Println("array3 = ", array3)

	array[2] = 35
	fmt.Println("array = ", array)

	array4 := [5]*int{0: new(int), 1: new(int)}
	fmt.Println("array4 = ", array4)

	*array4[0] = 10
	*array4[1] = 20
	fmt.Println("array4 = ", array4)

	// for item, value := range array4 {
	// 	if value != nil {
	// 		fmt.Println("item = ", *item)
	// 	}
	// }

	var array5 [5]string
	fmt.Println("array5 = ", array5)

	array6 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println("array6 = ", array6)

	var array7 [3]*string
	fmt.Println("array7 = ", array7)

	array8 := [3]*string{new(string), new(string), new(string)}

	*array8[0] = "Red"
	*array8[1] = "Blue"
	*array8[2] = "Green"

	array7 = array8

	fmt.Println("array8 = ", array8)

}

func MultiArray() {
	var array [4][2]int
	fmt.Println("array = ", array)

	array1 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println("array1 = ", array1)

	array2 := [4][2]int{1: {10, 11}, 3: {20, 21}}
	fmt.Println("array2 = ", array2)

	array3 := [4][2]int{1: {0: 20}, 3: {1: 4}}
	fmt.Println("array3 = ", array3)

	var array4 [2][2]int

	array4[0][0] = 10
	array4[0][1] = 20
	array4[1][0] = 30
	array4[1][1] = 40
	fmt.Println("array4 = ", array4)

}

func SliceArray() {
	slice := make([]string, 5)
	fmt.Println("slice = ", slice)

	slice1 := make([]string, 3, 5)
	fmt.Println("slice1 = ", slice1)

	slice2 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println("slice2 = ", slice2)

	slice3 := []int{10, 20, 30}
	fmt.Println("slice3 = ", slice3)

	slice4 := []string{99: ""}
	fmt.Println("slice4 = ", slice4)

	array := [3]int{10, 20, 30}
	fmt.Println("array = ", array)

	array1 := []int{10, 20, 30}
	fmt.Println("array1 = ", array1)

	var nil_slice []int
	fmt.Println("nil_slice = ", nil_slice)

	slice5 := make([]int, 0)
	fmt.Println("slice5  = ", slice5)

	slice6 := []int{}
	fmt.Println("slice6 = ", slice6)

	// length := slice6.Len()
	// fmt.Println("slice6 length = ", length)

	slice7 := []int{10, 20, 30, 40, 50}
	fmt.Println("slice7 = ", slice7)

	slice7[1] = 25
	fmt.Println("slice7 = ", slice7)

	newSlice := slice7[1:3]
	fmt.Println("newSlice = ", newSlice)
}
