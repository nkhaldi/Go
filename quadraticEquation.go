package main

import (
    "fmt"
    "math"
)


type xNum struct {
    Re, Im  float64
}


type Solution struct {
    x1, x2  xNum
    Success bool
    Comment string
}


func main() {
    a, b, c := getParameters()
    printEquation(a, b, c)

    sol := solveEquation(a, b, c)
    printSolution(sol)
}


func solveEquation(a, b, c float64) (Solution) {
    var x1, x2 float64

    if (a != 0) {
        var discr, sqrtD float64
        discr = b*b - 4*a*c
        if (discr >= 0) {
            sqrtD = math.Sqrt(discr)
            x1 = (-b - sqrtD) / (2 * a)
            x2 = (-b + sqrtD) / (2 * a)
        }
    } else if (b != 0) {
        x1 = -c / b
    }
    return x1, x2
}


func printEquation(a, b, c float64) {
    if (a != 0) {
        if (a == -1) {
            fmt.Print("-")
        } else if (a != 1) {
            fmt.Printf("%.2f", a)
        }
        fmt.Print("x^2")
    }
    if (b != 0) {
        if (a != 0) {
            if (b > 0) {
                fmt.Print(" + ")
            } else {
                fmt.Print(" - ")
            }
            if (b*b != 1) {
                fmt.Printf("%.2f*", b)
            }
        } else {
            if (b == -1) {
                fmt.Print("-")
            } else if (b != 1) {
                fmt.Printf("%.2f*", b)
            }
        }
        fmt.Print("x");
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


func printSolution(sol Solution) {
}


func getParameters() (float64, float64, float64) {
    var a, b, c float64
    fmt.Println("Enter parameters of equation ax^2 + bx + c = 0")

    fmt.Print("a = ");
	fmt.Scan(&a)
    fmt.Print("b = ");
	fmt.Scan(&b)
    fmt.Print("c = ");
	fmt.Scan(&c)

    return a, b, c
}
