package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {
	freqs := make(map[string]int)
	s := strings.Fields(text)
	for i := range s {
		if s[i][len(s[i])-1] == '.' || s[i][len(s[i])-1] == ',' {
			freqs[strings.ToLower(s[i][:len(s[i])-1])]++
		} else {
			freqs[strings.ToLower(s[i])]++
		}
	}
	return freqs
}

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read in DataFile as a string called data
	data, err := ioutil.ReadFile(DataFile)
	check(err)
	fmt.Printf("%#v", WordCount(string(data)))
	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
