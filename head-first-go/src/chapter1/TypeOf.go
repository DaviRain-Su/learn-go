package chapter1

import (
	"fmt"
	"reflect"
)

func TypeOf() {
	fmt.Println("\"42\" type is ", reflect.TypeOf(42))
	fmt.Println("\"3.1415\" type is", reflect.TypeOf(3.1415))
	fmt.Println("\"true\" type is", reflect.TypeOf(true))
	fmt.Println("\"Hello, World\" type is ", reflect.TypeOf("Hello, World!"))
}
