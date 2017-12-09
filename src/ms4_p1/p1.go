package main

import "fmt"
import "io/ioutil"
import "strings"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ByteArrToStrings(b []byte) [][]string {
    // Takes array of ASCII bytes, translates to integers
    contents := strings.TrimSpace(string(b))
    rows := strings.Split(contents, "\n")
    input := make([][]string, len(rows))
    for i, row := range rows {
        input[i] = strings.Fields(row)
    }
    return input
}

func removeDuplicatesUnordered(elements []string) []string {
    // Taken from https://www.dotnetperls.com/duplicates-go
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v:= range elements {
        encountered[elements[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key, _ := range encountered {
        result = append(result, key)
    }
    return result
}

func main() {
    fmt.Printf("Day 4, puzzle #1:\n")

    var valid_passes int = 0

    filebytes, err := ioutil.ReadFile("./inputs/ms4.txt")
    check(err)

    var input [][]string = ByteArrToStrings(filebytes)

    for _, row := range input {
        unordered := removeDuplicatesUnordered(row)
        if len(row) == len(unordered) {
            valid_passes++
        }
    }

    fmt.Printf("Acceptable passphrases: %d.\n", valid_passes)
}

