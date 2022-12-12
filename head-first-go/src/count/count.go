package count

import "fmt"

func Down(idx int) {
	fmt.Println("this idx is ", idx)
	if idx <= 1 {
		return
	} else {
		Down(idx - 1)
	}
}

func Greet(name string) {
	fmt.Println("hello,", name, "!")
	greet2(name)
	fmt.Println("getting ready to say bye ...")
	bye()
}

func greet2(name string) {
	fmt.Println("How are you,", name, "? ")
}

func bye() {
	fmt.Println("ok bye!")
}

func Fact(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * Fact(num-1)
	}
}

func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else if len(numbers) == 1 {
		return numbers[0]
	} else {
		return numbers[0] + Sum(numbers[1:])
	}
}

// Quick Sort
func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]

		var less []int
		var greater []int
		for _, v := range array[1:] {
			if v <= pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}

		var result []int
		result = append(result, QuickSort(less)...)
		result = append(result, pivot)
		result = append(result, QuickSort(greater)...)
		return result
	}
}

func DisplayGraph() {
	//graph :=
	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"aunj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}
	fmt.Printf("graph is %v\n", graph)
	fmt.Println("result is ", GraphSearch(graph, "you"))
}

func GraphSearch(graph map[string][]string, name string) bool {
	var searchQueue []string
	searchQueue = append(searchQueue, graph[name]...)
	var searched []string
	for len(searchQueue) != 0 {
		person := searchQueue[0]
		if !find(searched, person) {
			if personIsSeller(person) {
				fmt.Println(person, " is a mango seller!")
				return true
			} else {
				searchQueue = append(searchQueue, graph[person]...)
				searched = append(searched, person)
			}
		}
		searchQueue = searchQueue[1:]
	}
	return false
}

func find(list []string, value string) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}
	return false
}

func personIsSeller(name string) bool {
	return name[len(name)-1] == 'm'
}
