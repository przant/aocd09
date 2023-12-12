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
    pf, err := os.Open("input.txt")
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

    seq := make([]int64, 0)
    sSeq := make([][]int64, 0)

    for _, strNum := range strings.Split(strings.TrimSpace(history), " ") {
        num, _ := strconv.Atoi(strings.TrimSpace(strNum))
        seq = append(seq, int64(num))
    }

    sSeq = append(sSeq, seq)
    rightVal := seq[len(seq)-1]

    for i := 0; !allZeroes(sSeq[i]); i++ {
        sSeq = append(sSeq, nextSequence(sSeq[i]))
        if len(sSeq[i+1]) > 0 {
            rightVal += sSeq[i+1][len(sSeq[i+1])-1]
        }
    }
    leftVal := int64(0)
    for i := len(sSeq) - 1; i > 0; i-- {
        leftVal = -(leftVal - sSeq[i-1][0])
    }
    fmt.Println(leftVal, seq, rightVal)
    fmt.Println("***************************************************************************")
    return leftVal
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
        if row[i] != 0 {
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
