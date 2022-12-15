// echo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	fmt.Println("this program name: ", os.Args[0])
	for idx, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println("idx is ", idx, ", value is ", arg)
	}
	fmt.Println(s)
	// echo 的第二种方式的实现
	// fmt.Println(strings.Join(os.Args[1:], " "))
	// todo
	// 性能测试测试两个版本的实现
}
