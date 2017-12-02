package main

import "fmt"
import "io/ioutil"
import "sort"
import "strconv"
import "strings"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ByteArrToIntArr(b []byte) [][]int {
    // Takes array of ASCII bytes, translates to integers
    var err error
    contents := strings.TrimSpace(string(b))
    rows := strings.Split(contents, "\n")
    spreadsheet := make([][]int, len(rows))
    for i, row := range rows {
        s := strings.Fields(row)
        spreadsheet[i] = make([]int, len(s))
        for j, field := range s {
            spreadsheet[i][j], err = strconv.Atoi(field)
            check(err)
        }
    }
    return spreadsheet
}

func main() {
    fmt.Printf("Day 2, puzzle #2:\n")

    checksum := 0

    filebytes, err := ioutil.ReadFile("./inputs/ms2.txt")
    check(err)

    spreadsheet := ByteArrToIntArr(filebytes)
    for _, row := range spreadsheet {
        sort.Ints(row)
        for _, n := range row {
            for _, m := range row {
                if n % m == 0 && n != m {
                    checksum += n/m
                    break
                }
            }
        }
    }

    fmt.Printf("Final checksum: %d.\n", checksum)
}

