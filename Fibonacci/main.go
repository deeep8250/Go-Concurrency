package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonacci(wg *sync.WaitGroup, n int, channel chan int) {

	defer wg.Done()

	a, b := 0, 1

	channel <- a
	channel <- b
	for i := 0; i < n-2; i++ {
		c := a + b
		channel <- c
		time.Sleep(time.Second * 1)
		a = b
		b = c

	}
	defer close(channel)

}

func main() {

	channel := make(chan int)

	n := 6
	var wg sync.WaitGroup

	wg.Add(1)
	go Fibonacci(&wg, n, channel)

	go func() {
		for i := range channel {
			fmt.Println(i)

		}

	}()

	wg.Wait()

}
