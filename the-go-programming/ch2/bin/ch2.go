package main

import "fmt"

// # 2.1 命名
// 如果一个名字是大写字母开头的（译注：必须是在函数外部定义的包级名字；
// 包级函数名本身也是包级名字），那么它将是导出的，也就是说可以被外部的包访问，
// 例如fmt包的Printf函数就是导出的，可以在fmt包外部访问。包本身的名字一般总是用小写字母。
//
// 在习惯上，Go语言程序员推荐使用 驼峰式 命名，当名字由几个单词组成时优先使用大小写分隔，而不是优先用下划线分隔。
//
// 在包一级声明语句声明的名字可在整个包对应的每个源文件中访问，而不是仅仅在其声明语句所在的源文件中访问。
// 局部声明的名字就只能在函数内部很小的范围被访问。

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
