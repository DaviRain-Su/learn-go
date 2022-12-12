package chapter8

import "fmt"

type MyStruct struct {
	name string
	age  int
}

type Table struct {
	number float64
	word   string
	toggle bool
}

func DisplayStruct() {
	fmt.Printf("%#v\n", MyStruct{})
	var myStruct MyStruct
	myStruct.age = 13
	myStruct.name = "davirain"
	fmt.Printf("%#v\n", myStruct)

	var subsctriber Subscriber
	subsctriber.Name = "davirian"
	subsctriber.Rate = 23.0
	subsctriber.Active = true
	fmt.Printf("%#v\n", subsctriber)

	subsctriber2 := Subscriber{Name: "davirian", Rate: 23.0}
	fmt.Printf("%#v\n", subsctriber2)

	example := Example{"davirain"}
	fmt.Println("example is ", example.Text())
}
