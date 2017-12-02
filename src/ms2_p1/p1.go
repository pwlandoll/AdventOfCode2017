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
    fmt.Printf("Day 1, puzzle #1:\n")

    checksum := 0
    difference := 0

    filebytes, err := ioutil.ReadFile("./inputs/ms2.txt")
    check(err)

    spreadsheet := ByteArrToIntArr(filebytes)
    fmt.Println(spreadsheet)
    for _, row := range spreadsheet {
        sort.Ints(row)
        difference = row[len(row)-1] - row[0]
        checksum += difference
    }

    fmt.Printf("Final checksum: %d.\n", checksum)
}

