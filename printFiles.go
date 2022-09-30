/*
   Написать функцию, которая рекурсивно печатает файлы
   из указанной директории
*/

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	printAllFiles(".")
}

func printAllFiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, file := range files {
		filename := filepath.Join(path, file.Name())
		fmt.Println(filename)
		if file.IsDir() {
			printAllFiles(filename)
		}
	}
}
