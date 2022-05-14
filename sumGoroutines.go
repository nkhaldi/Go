/*
   Программа должна вычислить и вывести сумму всех четных чисел,
   в заданном диапазоне плюс сумму их квадратов, используя горутины,
   и вывести результат
*/

package main

import "fmt"

func evenSum(x1, x2 int, ch chan int) {
	result := 0
	for num := x1; num <= x2; num++ {
		if num%2 == 0 {
			result += num
		}
	}
	ch <- result
}

func squareSum(x1, x2 int, ch chan int) {
	result := 0
	for num := x1; num <= x2; num++ {
		if num%2 == 0 {
			result += num * num
		}
	}
	ch <- result
}

func main() {
	evnCh := make(chan int)
	sqrCh := make(chan int)

	go evenSum(0, 100, evnCh)
	go squareSum(0, 100, sqrCh)

	evnSum := <-evnCh
	sqrSum := <-sqrCh
	fmt.Println("Even sum =", evnSum)
	fmt.Println("Square sum =", sqrSum)
	fmt.Println("Total sum =", evnSum+sqrSum)
}
