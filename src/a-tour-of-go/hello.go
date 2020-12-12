package main

//import (
//	"fmt"
//	"math/rand"
//)

import "fmt"

// 导出名
// 如果一个名字大写字母组成，那么他就是已导出的，ex， math.Pi
// math.Pi
func main()  {
	//fmt.Println("Hello, Wold!")
	//fmt.Println("My favorite number is ", rand.Intn(10))
	//fmt.Printf("Now you have %g problems.\n\n", math.Sqrt(9))
	//fmt.Println(math.Pi)
	//fmt.Println(add(42, 13))
	//a, b := swap("hello", "world")
	//fmt.Printf("a = %g, b = %g\n\n", a, b)
	//fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
}

// var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后
// var 语句可以出现在包或者函数级别
var c, python, java bool


// Go 的返回值可以被命名， 他们会被视作定义在函数顶部的变量
// 返回值的名称应当具有一定的意义， 它可以被作为文档使用
// 没有参数的retun语句返回已命名的返回值， 也就是直接返回
// 直接返回语句应当用在短函数中，在长函数中它们会影响代码的可读性。
func split(sum int) (x, y int)  {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

/// 函数可以没有参数或接受多个参数
//func add(x int, y int) int {
func add(x, y int) int { // 省略格式
	return x + y
}
