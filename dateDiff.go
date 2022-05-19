/*На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	rd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}

	dates := strings.Split(strings.Trim(rd, "\n"), ",")
	date1, err := time.Parse("02.01.2006 15:04:05", dates[0])
	if err != nil {
		panic(err)
	}
	date2, err := time.Parse("02.01.2006 15:04:05", dates[1])
	if err != nil {
		panic(err)
	}

	if date2.Before(date1) {
		date1, date2 = date2, date1
	}

	diff := date2.Sub(date1)
	fmt.Println(diff)
}
