package main

import "fmt"
import "math"

func spiral_coordinates(n int) (int, int) {
    /* Time to do some math!
     * _informed_ by the stackexchange answer below:
     * https://math.stackexchange.com/q/163101
     */
    k := int(math.Ceil((math.Sqrt(float64(n))-1)/2))
    t := 2*k + 1
    m := t*t
    t -= 1
    if n >= m-t {
        return k-(m-n), -k
    } else {
        m = m-t
    }
    if n >= m-t {
        return -k, -k+(m-n)
    } else {
        m = m-t
    }
    if n >= m-t {
        return -k+(m-n), k
    } else {
        return k, k-(m-n-t)
    }
}

func main() {
    fmt.Printf("Day 3, puzzle #1:\n")
    input := 347991
    distance := 0

    x, y := spiral_coordinates(input)

    distance = int(math.Abs(float64(x)) + math.Abs(float64(y)))

    fmt.Printf("Distance: %d.\n", distance)
}

