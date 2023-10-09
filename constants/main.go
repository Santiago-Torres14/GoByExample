package main

import (
    "fmt"
    "math"
)

const s = "constant"

func main() {
    fmt.Printf("%v\n", s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Printf("%v\n", d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}

