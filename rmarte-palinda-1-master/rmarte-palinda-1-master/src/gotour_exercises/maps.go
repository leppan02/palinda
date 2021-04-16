package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	a := strings.Fields(s)
	b := map[string]int{}
	for i := 0; i < len(a); i++ {
		b[a[i]]++
	}
	return b
}

func main() {
	wc.Test(WordCount)
}
