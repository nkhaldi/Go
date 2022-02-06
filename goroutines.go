package main

import (
    "fmt"
    "time"
)


func goPrinter(x1, x2 int, ch chan bool) {
    for i := x1; i <= x2; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(i)
    }
    ch <- true
}


func main() {
    ch := make(chan bool)
    fmt.Println("Begin")
    go goPrinter(0, 5, ch)
    go goPrinter(6, 10, ch)
    fmt.Println("Fin")
    <-ch
    <-ch
    fmt.Println("End")
}
