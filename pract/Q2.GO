
package main

import (
    "fmt"
    "strconv"
)

func main() {
    v, err := strconv.Atoi("A")

    if err != nil {
        fmt.Println("bad input:", err)
        return
    }

    fmt.Println(v)
}
