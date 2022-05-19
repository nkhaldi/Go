/*
   На стандартный ввод подаются данные о студентах университетской группы
   в формате JSON. В сведениях о каждом студенте содержится информация
   о полученных им оценках (Rating). Требуется прочитать данные
   и рассчитать среднее количество оценок, полученное студентами группы.
   Ответ на задачу требуется записать на стандартный вывод в формате JSON
   в следующей форме:
   { "Average": 14.1 }
*/

package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Group struct {
	ID       int
	Year     int
	Number   string
	Students []Student
}

type avgStruct struct {
	Average float64
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var gr Group
	err = json.Unmarshal(data, &gr)
	if err != nil {
		panic(err)
	}

	var cnt int = 0
	for _, stud := range gr.Students {
		cnt += len(stud.Rating)
	}
	avg := float64(cnt) / float64(len(gr.Students))
	ansStruct := avgStruct{
		Average: avg,
	}

	answer, err := json.MarshalIndent(ansStruct, "", "    ")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(answer)
}
