package main

import "fmt"

func main() {
	fmt.Println("Begin")
	for i := 1; i <= 5; i++ {
		defer fmt.Println("Text", i)
	}
	fmt.Println("End")
}
