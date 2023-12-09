package main

import (
    "bufio"
    "log"
    "os"
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
}
