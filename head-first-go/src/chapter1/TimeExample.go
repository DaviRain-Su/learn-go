package chapter1

import (
	"fmt"
	"time"
)

func GetCurrentTime() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println("Now is ", now, "and this year is ", year)
}
