package main

import (
	"fmt"
	"headfirstgo/src/chapter5/datafile"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	numbers, err := datafile.GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("sum : %0.2f\n", sum)

	//err = Socialize()
	//if err != nil {
	//	log.Fatal(err)
	//}

	defer reportPanic()
	//count(1, 3)
	panic("Some other issue")
	scanDirectory("...")
	//defer calmDown()
	//err = fmt.Errorf("there's an error")
	//panic(err)
}

func Socialize() error {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("I don't want to talk.")
	fmt.Println("Nice weather, eth?")
	return nil
}

func count(start int, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
	fmt.Printf("Returning from count(%d, %d) call\n", start, end)
}

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(err)
	}
}
