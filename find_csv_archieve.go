/*
   В тестовом архиве "task.zip" содержится набор папок и файлов.
   Один из этих файлов является файлом с данными в формате CSV,
   прочие же файлы структурированных данных не содержат.
   Требуется найти и прочитать этот файл со структурированными данными
   (это CSV-таблица 10х10, разделителем является запятая),
   а в качестве ответа необходимо указать число,
   находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
*/

package main

import (
	"fmt"
	"archive/zip"
	"encoding/csv"
)

func main() {
	archieve, err := zip.OpenReader("tests/task.zip")
	defer archieve.Close()
	if err != nil {
		panic(err)
	}

	for _, file := range archieve.File {
		fh, err := file.Open()
		defer fh.Close()
		if err != nil {
			panic(err)
		}

		lines, err := csv.NewReader(fh).ReadAll()
		if err != nil {
			panic(err)
		}
		if len(lines) == 10 && len(lines[4]) == 10 {
			fmt.Println("The answer is", lines[4][2])
			break
		}
	}
}
