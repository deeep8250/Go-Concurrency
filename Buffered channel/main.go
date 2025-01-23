// QUESTION -> Write a program where one goroutine adds values (1 to 5) to the channel, and another goroutine reads these values.
package main

import (
	"fmt"
	"sync"
)

func Sender(wg *sync.WaitGroup, n int, channel chan int) {
	defer wg.Done()

	for i := 1; i <= n; i++ {

		channel <- i

	}
	close(channel)

}

func main() {

	channel := make(chan int, 5)
	var wg sync.WaitGroup
	n := 5

	wg.Add(2)
	go Sender(&wg, n, channel)
	go func() {
		defer wg.Done()
		for i := range channel {
			fmt.Println(i)
		}

	}()

	wg.Wait()

}
