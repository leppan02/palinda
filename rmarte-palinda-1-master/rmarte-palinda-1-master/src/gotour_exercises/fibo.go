package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	a, b := 0, 1
	return func() int{
		a, b, c := b, a+b, a
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
