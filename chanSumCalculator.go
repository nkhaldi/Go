/*
   Вам необходимо написать функцию calculator следующего вида:
   func calculator(args <-chan int, done <-chan struct{}) <-chan int
   Через канал args функция получит ряд чисел,
   а через канал done - сигнал о необходимости завершить работу.
   Когда сигнал о завершении работы будет получен, функция должна
   в выходной канал отправить сумму полученных чисел.
   Функция calculator должна быть неблокирующей, сразу возвращая управление.
   Выходной канал должен быть закрыт после выполнения всех операций,
   если вы этого не сделаете, то превысите предельное время работы.
*/

package main

import "fmt"

func main() {
	args := make(chan int)
	done := make(chan struct{})

	go func() {
		defer close(args)
		defer close(done)

		fmt.Println("Enter some numbers and finish the list with '0'")
		for {
			var num int
			fmt.Scan(&num)
			if num != 0 {
				args <- num
			} else {
				break
			}
		}
		done <- struct{}{}
	}()

	res := <-calculator(args, done)
	fmt.Printf("The result is %d\n", res)
}

func calculator(args <-chan int, done <-chan struct{}) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		var sum int = 0
		for {
			select {
			case value := <-args:
				sum += value
			case <-done:
				result <- sum
				return
			}
		}
	}()
	return result
}
