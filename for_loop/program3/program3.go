// Write a program to print the output like
// 1
// 22
// 333
// 4444
// 55555
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print(i)
        }
        fmt.Println()
    }
}

