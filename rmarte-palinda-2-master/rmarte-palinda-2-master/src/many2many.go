// Stefan Nilsson 2013-03-13

// This is a testbed to help you understand channels better.

// * What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
// **** kommer stänga kanalen för andra genom att att det är samma kanl

// * What happens if you remove the statement close(ch) completely?
// **** Nothing

// * Can you be sure that all strings are printed before the program stops?
// **** NO
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 4 // double speed

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgc := new(sync.WaitGroup)
	wgp.Add(producers)
	wgc.Add(consumers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch, wgc)
	}

	wgp.Wait() // Wait for all producers to finish.
	wgc.Wait() // Wait for all consumers to finish.
	close(ch)  // om du byter kommer den skicka på en stängd kanal
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
	wg.Done()
}

// RandomSleep waits for x ms, where x is a random number, 0 â‰¤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
