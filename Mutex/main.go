// Question -> Problem: Write a program where multiple goroutines increment a shared counter variable.

package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	routine := 10
	var wg sync.WaitGroup
	var mutex sync.Mutex
	channel := make(chan int)

	for i := 0; i < routine; i++ {
		wg.Add(1)
		go SharedLocation(&wg, channel, &counter, &mutex)
	}

	go func() {
		defer close(channel)
		for i := range channel {
			fmt.Println(i)
		}

	}()

	wg.Wait()
}

func SharedLocation(wg *sync.WaitGroup, channel chan int, counter *int, mutex *sync.Mutex) {

	defer wg.Done()
	mutex.Lock()
	for i := 0; i < 10; i++ {

		*counter += 1
		channel <- *counter

	}
	mutex.Unlock()

}
