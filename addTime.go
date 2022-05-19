/*
   На стандартный ввод подаются данные о продолжительности периода
   в формате "12 мин. 13 сек.". Кроме того, дана дата в формате
   Unix-Time: 1589570165 в виде константы типа int64
   (наносекунды для целей преобразования равны 0).
   Требуется считать данные о продолжительности периода,
   преобразовать их в тип Duration, а затем вывести дату и время
   в формате UnixDate получившиеся при добавлении периода к начальноый дате.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const startTime int64 = 1589570165
const UnixDate = "Mon Jan _2 15:04:05 MST 2006"

func main() {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	arr := strings.Split(strings.Trim(input, "\n"), " ")

	durationFormat := arr[0] + "m" + arr[2] + "s"
	dur, err := time.ParseDuration(durationFormat)
	if err != nil {
		panic(err)
	}

	startDate := time.Unix(startTime, 0).UTC()
	endDate := startDate.Add(dur)

	fmt.Println(endDate.Format(UnixDate))
}
