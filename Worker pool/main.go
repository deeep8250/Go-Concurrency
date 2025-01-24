package main

import (
	"fmt"
	"sync"
)

func Worker(wg *sync.WaitGroup, start int, end int, channel chan int) {
	defer wg.Done()

	for i := start; i <= end; i++ {
		channel <- i * i
	}
}

func main() {
	n := 5
	routine := 5
	chunk := n / routine
	remainder := n % routine

	var wg sync.WaitGroup
	channel := make(chan int)

	start := 1
	end := chunk

	for i := 0; i < routine; i++ {
		wg.Add(1)
		go Worker(&wg, start, end, channel)
		start = end + 1
		end += chunk
		if i < remainder {
			end++
			remainder--
		}

	} //go routine for

	go func() {
		defer close(channel)
		for i := range channel {
			fmt.Println(i)
		}
	}()
	wg.Wait()

}
