package chapter7

import (
	"fmt"
	"headfirstgo/src/chapter5/datafile"
	"log"
)

func Count() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line] += 1
	}
	for key, value := range counts {
		fmt.Printf("Votes for %s: %d\n", key, value)
	}
}
