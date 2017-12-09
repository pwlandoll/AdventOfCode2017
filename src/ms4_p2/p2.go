package main

import "fmt"
import "io/ioutil"
import "sort"
import "strings"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func SortString(w string) string {
    // Taken from https://stackoverflow.com/a/22689818
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func ByteArrToStrings(b []byte) [][]string {
    // Takes array of ASCII bytes, translates to integers
    contents := strings.TrimSpace(string(b))
    rows := strings.Split(contents, "\n")
    input := make([][]string, len(rows))
    for i, row := range rows {
        row_as_strings := strings.Fields(row)
        input[i] = make([]string, len(row_as_strings))
        for j, s := range row_as_strings {
            input[i][j] = SortString(s)
        }
    }
    return input
}

func removeDuplicatesUnordered(elements []string) []string {
    // Taken from https://www.dotnetperls.com/duplicates-go and modified
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v := range elements {
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
    fmt.Printf("Day 4, puzzle #2:\n")

    var valid_passes int

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

