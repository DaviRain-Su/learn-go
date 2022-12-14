package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
	}

	myChannel := make(chan Page)
	for _, url := range urls {
		go responseSize(url, myChannel)
	}
	for i := 0; i < len(urls); i += 1 {
		pages := <-myChannel
		fmt.Printf("%s: %d\n", pages.URL, pages.Size)
	}

	//go a()
	//go b()
	//time.Sleep(5 * time.Second)
	//fmt.Println("end main()")
	//
	//myChannel := make(chan string)
	//fmt.Println("my channel ", myChannel)
	//go greeting(myChannel)
	//receivedValue := <-myChannel
	//fmt.Println("receive value is ", receivedValue)
	//channel1 := make(chan string)
	//channel2 := make(chan string)
	//go abc(channel1)
	//go def(channel2)
	//fmt.Println(<-channel1)
	//fmt.Println(<-channel2)
	//fmt.Println(<-channel1)
	//fmt.Println(<-channel2)
	//fmt.Println(<-channel1)
	//fmt.Println(<-channel2)
	//fmt.Println()

	//myChannel := make(chan string)
	//go send(myChannel)
	//reportNap("receiving goroutine", 5)
	//fmt.Println(<-myChannel)
	//fmt.Println(<-myChannel)

}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

func a() {
	for i := 0; i < 50; i += 1 {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i += 1 {
		fmt.Print("b")
	}
}

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i += 1 {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "Wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}
