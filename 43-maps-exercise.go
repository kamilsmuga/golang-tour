package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    count := make(map[string]int)
    words := strings.Fields(s)
    for _,str := range(words) {
        count[str] += 1
    }
    return count
}

func main() {
    wc.Test(WordCount)
}
