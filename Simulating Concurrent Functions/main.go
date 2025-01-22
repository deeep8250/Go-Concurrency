package main

import (
	"fmt"
	"sync"
	"time"
)

func goNumPrint(wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		fmt.Println("for", n, " : ", i)
		time.Sleep(time.Second * 1)
	}

}

func main() {
	var wg sync.WaitGroup
	db := []int{5, 10, 15}

	for _, dbc := range db {
		wg.Add(1)
		go goNumPrint(&wg, dbc)
	}
	wg.Wait()

}
