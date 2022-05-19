package main

import (
	"fmt"
	"strconv"
	"time"
)

func tester(num int, ch chan string) {
	res := "Text #" + strconv.Itoa(num)
	time.Sleep(100 * time.Millisecond)
	ch <- res
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	for i := 0; i < 20; i++ {
		go tester(1, ch1)
		go tester(2, ch2)
		select {
		case test1 := <-ch1:
			fmt.Println(test1)
		case test2 := <-ch2:
			fmt.Println(test2)
		default:
			fmt.Println("Ничего не доступно")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
