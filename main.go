package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    pf, err := os.Open("example.txt")
    if err != nil {
        log.Fatalf("while opening file %q: %s", pf.Name(), err)
    }
    defer pf.Close()

    scnr := bufio.NewScanner(pf)

    lines := make([]string, 0)
    for scnr.Scan() {
        lines = append(lines, scnr.Text())
    }

    total := int64(0)
    for _, history := range lines {
        total += predict(history)
    }
    fmt.Println(total)
}

func predict(history string) int64 {
    sequence := make([]int64, 0)

    for _, strNum := range strings.Split(strings.TrimSpace(history), " ") {
        num, _ := strconv.Atoi(strings.TrimSpace(strNum))
        sequence = append(sequence, int64(num))
    }

    result := sequence[len(sequence)-1]
    for !allZeroes(sequence) {
        sequence = nextSequence(sequence)
        if len(sequence) > 0 {
            result += sequence[len(sequence)-1]
        }
    }
    return result
}

func nextSequence(seq []int64) []int64 {
    ns := make([]int64, 0)

    for i := 0; i < len(seq)-1; i++ {
        ns = append(ns, seq[i+1]-seq[i])
    }
    return ns
}

func allZeroes(row []int64) bool {
    for i := 0; i < len(row); i++ {
        if row[i] > 0 {
            return false
        }
    }

    return true
}

func atoi(str string) uint {
    num, err := strconv.Atoi(str)
    if err != nil {
        log.Printf("while parsing string %q into a num: %s", str, err)
        return 0
    }

    return uint(num)
}
