package main

import "fmt"

func main() {
    for {
    a, b, c := getParameters()
    printEquation(a, b, c)
}
    //x1, x2 := solveEquation(a, b, c)
    //fmt.Printf("x1 = %.2f, x2 = %.2f\n", x1, x2)
}


func solveEquation(a, b, c float32) (float32, float32) {
    var x1, x2 float32

    x1 = 123.45
    x2 = 321.54
    return x1, x2
}


func printEquation(a, b, c float32) {
    if (a != 0) {
        if (a == -1) {
            fmt.Print("-")
        } else if (a != 1) {
            fmt.Print(a)
        }
        fmt.Print("*x^2")
    }
    if (b != 0) {
        if (a != 0) {
            if (b > 0) {
                fmt.Print(" + ")
            } else {
                fmt.Print(" - ")
            }
            if (b*b != 1) {
                fmt.Printf("%.2f", b)
            }
        } else {
            if (b == -1) {
                fmt.Print("-")
            } else if (b != 1) {
                fmt.Printf("%.2f", b)
            }
        }
        fmt.Print("*x");
    }
    if (c != 0) {
        if (a != 0 || b != 0) {
            if (c > 0) {
                fmt.Printf(" + %.2f", c)
            } else {
                fmt.Printf(" - %.2f", -c)
            }
        } else {
            fmt.Printf("%.2f", c)
        }
    }
    if (a == 0 && b == 0 && c == 0) {
        fmt.Print("0")
    }
    fmt.Println(" = 0")
}


func getParameters() (float32, float32, float32) {
    var a, b, c float32
    fmt.Println("Enter parameters of equation:")
    fmt.Println("ax^2 + bx + c = 0")
    fmt.Print("a = ");
	fmt.Scan(&a)
    fmt.Print("b = ");
	fmt.Scan(&b)
    fmt.Print("c = ");
	fmt.Scan(&c)
    return a, b, c
}
