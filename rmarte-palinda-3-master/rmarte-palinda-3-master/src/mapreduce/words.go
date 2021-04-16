package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

type element struct {
	key string
	val int
}

var s []string

func calc(start int, stop int, ch chan element, wg *sync.WaitGroup) {
	freqs := make(map[string]int)
	for i := start; i < stop; i++ {
		if s[i][len(s[i])-1] == '.' || s[i][len(s[i])-1] == ',' {
			freqs[strings.ToLower(s[i][:len(s[i])-1])]++
		} else {
			freqs[strings.ToLower(s[i])]++
		}
	}
	for key, val := range freqs {
		ch <- element{key, val}
	}
	wg.Done()
}

func WordCount(text string) map[string]int {
	var wg sync.WaitGroup
	freqs := make(map[string]int)
	s = strings.Fields(text)
	threads := 6
	wg.Add(threads)
	l := len(s) / (threads)
	ch := make(chan element)
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := 0; i < threads-1; i++ {
		go calc(i*l, (i+1)*l, ch, &wg)
	}
	go calc((threads-1)*l, len(s), ch, &wg)
	for i := range ch {
		freqs[i.key] += i.val
	}
	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, _ := ioutil.ReadFile(DataFile)
	fmt.Printf("%#v", WordCount(string(data)))
	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
