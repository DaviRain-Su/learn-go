package chapter5

import (
	"fmt"
	"headfirstgo/src/chapter5/datafile"
	"log"
	"time"
)

func ArrayExample() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	//notes[0] = "do"
	//notes[1] = "re"
	//notes[2] = "mi"
	//fmt.Println(notes)
	//fmt.Printf("%#v\n", notes)
	for index, value := range notes {
		fmt.Println("index = ", index, ", value = ", value)
	}

	primes := [5]int{2, 3, 5, 7, 11}
	//primes[0] = 2
	//primes[1] = 3
	//fmt.Println(primes)
	//fmt.Printf("%#v\n", primes)
	for index, value := range primes {
		fmt.Println("index = ", index, ", value = ", value)
	}

	var dates [3]time.Time
	dates[0] = time.Now()
	dates[1] = time.Now()
	dates[2] = time.Now()
	//fmt.Println(dates)
	//fmt.Printf("%#v\n", dates)
	for index, value := range dates {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//number := [3]float64{2, 3, 4}
	//CalcArrayAverage(number)

	number, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	CalcArrayAverage(number)
}

func CalcArrayAverage(array []float64) float64 {
	var sum float64 = 0
	for _, number := range array {
		sum += number
	}
	fmt.Println("sum is ", sum)
	fmt.Println("average is ", sum/float64(len(array)))
	return sum / float64(len(array))
}
