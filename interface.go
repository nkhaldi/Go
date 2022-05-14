package main

import (
	"os"
	"fmt"
	"encoding/json"
)

func main() {
	// readTask возвращзает значения для операндов и оператора
	// в формате пустого интерфейса
	value1, value2, operation := readTask()

	v1, ok1 := value1.(float64)
	if !ok1 {
		fmt.Print("value=", value1, ": ")
		fmt.Printf("%T", value1)
		return
	}

	v2, ok2 := value2.(float64)
	if !ok2 {
		fmt.Print("value=", value2, ": ")
		fmt.Printf("%T", value2)
		return
	}

	switch op := operation.(type) {
	case string:
		if op == "+" || op == "-" || op == "*" || op == "/" {
			fmt.Printf("%.4f", calc(v1, v2, op))
		} else {
			fmt.Println("неизвестная операция")
			return
		}
	default:
		fmt.Println("неизвестная операция")
		return
	}

}

func calc(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		} else {
			return 0.0
		}
	default:
		return 0.0
	}
}
