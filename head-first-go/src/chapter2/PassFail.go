package chapter2

import (
	"fmt"
	//"headfirstgo/src/keyboard"
	"github.com/headfirstgo/keyboard"
	"log"
	"os"
)

// pass_fail reports whether a grade is passing or failing.
func PassFail() {
	fmt.Print("Enter a grade: ")

	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade == 100 {
		status = "Perfect!"
	} else if grade >= 60 {
		status = "Passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of ", grade, "is", status)
}

// get main.go file size
func FileSize() {
	fileInfo, err := os.Stat("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()
	fmt.Println("main.go file size is ", fileSize)
}

func blockSpace() {
	var a = "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
	}
	fmt.Println(a)
	fmt.Println(b)
}
