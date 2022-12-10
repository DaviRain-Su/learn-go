package chapter5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(file_name string) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
