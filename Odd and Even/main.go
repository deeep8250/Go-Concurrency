package main

import (
	"fmt"
	"math"
	"sync"
)

func EvenOdd(wg *sync.WaitGroup, start int, end int, channel chan string) {
	defer wg.Done()

	for i := start; i <= end; i++ {

		sqrt := math.Sqrt(float64(i))
		isprime := true

		for j := 2; j <= int(sqrt); j++ {
			if i%j == 0 {
				isprime = false
				break
			}
		}

		if isprime == true {
			channel <- fmt.Sprintf("%d is a prime number", i)

		} else {
			fmt.Println(i, " is not a prime number")

		}

	}

}

func main() {

	channel := make(chan string)
	n := 20
	goroutine := 4
	chunk := n / (goroutine - 1)
	remainder := n % (goroutine - 1)
	var wg sync.WaitGroup

	start := 2
	end := chunk

	if remainder != 0 {
		wg.Add(1)
		x := n - remainder
		y := n
		//fmt.Println("chunk is from in first if ", x, " to ", y)
		go EvenOdd(&wg, x+1, y, channel)
	}

	for i := 1; i < goroutine; i++ {

		wg.Add(1)
		//fmt.Println("chunk is from ", start, " to ", end)
		go EvenOdd(&wg, start, end, channel)
		start = end + 1
		end += chunk

	}

	go func() {
		defer close(channel)
		for i := range channel {
			fmt.Println(i)
		}

	}()

	wg.Wait()

}
