package chapter7

import (
	"fmt"
)

func DisplayMap() {
	var myMap map[string]float64
	myMap = make(map[string]float64)
	myMap["gold"] = 1
	myMap["silver"] = 2
	myMap["bronze"] = 3
	fmt.Println("myMap is ", myMap)

	myMap2 := map[string]float64{"a": 1.23, "b": 3.4}
	fmt.Println("myMap2 is ", myMap2)

	emptyMap := map[string]float64{}
	fmt.Println("emptyMap is ", emptyMap)
	status("Carl")

	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)

	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("priome: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("priome: %v, ok: %v\n", prime, ok)

}

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}
