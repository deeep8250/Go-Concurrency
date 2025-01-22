package main

import (
	"fmt"
	"sync"
	"time"
)

func Evengoroute(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is even task 1")
		}
		time.Sleep(1 * time.Second)
	}

}

func Oddgoroute(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i, " is odd task 2")
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go Evengoroute(&wg)

	go Oddgoroute(&wg)

	wg.Wait()

}
