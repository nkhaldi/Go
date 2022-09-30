/*
	Фунеция генерации чисел Фиббоначи, построенная на замыканиях
*/

package main

import "fmt"

func main() {
	f := fib()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fib() func() int {
	x1, x2 := 0, 1

	return func() int {
		x1, x2 = x2, x1+x2
		return x1
	}
}
