package main

//import "bytes"
import "fmt"
import "io/ioutil"
import "strconv"
import "strings"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ByteArrToIntArr(b []byte) []int {
    // Takes array of ASCII bytes, translates to integers
    contents := strings.TrimSpace(string(b))
    contents_array := strings.Split(contents, "")
    int_arr := make([]int, len(contents_array))
    var err error
    for i := range int_arr {
        int_arr[i], err = strconv.Atoi(contents_array[i])
        check(err)
    }
    return int_arr
}

func main() {
    fmt.Printf("Day 1, puzzle #1:\n")

    sum := 0

    filebytes, err := ioutil.ReadFile("./ms1/input.txt")
    check(err)

    int_arr := ByteArrToIntArr(filebytes)
    array_length := len(int_arr)

    for i, v := range int_arr {
        if v == int_arr[(i+1) % array_length] {
            sum += v
        }
    }

    fmt.Printf("Final sum: %d.", sum)
}

