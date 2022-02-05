package main

import (
    "fmt"
    "time"
)


func go_printer(x1, x2 int) {
    for i := x1; i <= x2; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(i)
    }
}


func main() {
    fmt.Println("Begin")
    go go_printer(0, 5)
    go go_printer(6, 10)
    fmt.Println("Fin")
    time.Sleep(1000 * time.Millisecond)
    fmt.Println("End")
}
