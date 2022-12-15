package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello, world!")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()] += 1
	}

	// Note: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
