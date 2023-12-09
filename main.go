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

    total := 0
    for _, line := range lines {
        total += history(line)
    }
    fmt.Println(total)
}

func history(line string) int {
    sSeqs := make([][]int, 0)
    nums := strings.Fields(line)

    seq1 := make([]int, 0)
    for _, num := range nums {
        seq1 = append(seq1, atoi(num))
    }

    sSeqs = append(sSeqs, seq1)

    sum := sSeqs[0][len(sSeqs[0])-1]
    for i := 0; i < len(sSeqs[0]); i++ {
        zeroDiff := true
        seqN := make([]int, 0)
        length := len(sSeqs[0])
        for j := 0; j < length-(i+1); j++ {
            diff := sSeqs[i][j+1] - sSeqs[i][j]
            seqN = append(seqN, diff)
            if diff > 0 {
                zeroDiff = false
            }
        }
        sum += seqN[len(seqN)-1]
        sSeqs = append(sSeqs, seqN)
        if zeroDiff {
            break
        }
    }
    return sum
}

func atoi(str string) int {
    num, err := strconv.Atoi(str)
    if err != nil {
        log.Printf("while parsing string %q into a num: %s", str, err)
        return 0
    }

    return num
}
