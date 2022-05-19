/*
   В тестовом файле "task.data", содержится длинный ряд чисел,
   разделенных символом ";". Требуется найти, на какой позиции
   находится число 0 и указать её в качестве ответа.
   Требуется вывести именно позицию числа, а не индекс.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("tests/task.data")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(file)
	var pos int = 1
	for {
		str, err := r.ReadString(';')
		if err != nil {
			panic(err)
		}
		if str == "0;" {
			fmt.Println(pos)
			break
		}
		pos++
	}
}
