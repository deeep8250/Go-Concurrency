// Create three different tasks (simulated by goroutines) that do the following:

// Task 1: Simulate a task that takes 2 seconds to complete.
// Task 2: Simulate a task that takes 3 seconds to complete.
// Task 3: Simulate a task that takes 1 second to complete.
// Each task should send a message through its own channel when it's done.

// The main program should use select to listen for the results from the tasks and print out the messages in the order they finish. If all tasks finish at the same time, select should randomly pick one task's message.

package main

import (
	"fmt"
	"sync"
	"time"
)

func chA(channelA chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	channelA <- "hello from channel 1"

}
func chB(channelB chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	channelB <- "hello from channel 2"

}
func chC(channelC chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	channelC <- "hello from channel 3"

}

func main() {
	channelA := make(chan string)
	channelB := make(chan string)
	channelC := make(chan string)
	var wg sync.WaitGroup

	wg.Add(3)
	go chA(channelA, &wg)
	go chB(channelB, &wg)
	go chC(channelC, &wg)

	for i := 0; i < 3; i++ {
		select {
		case msg, ok := <-channelA:
			if ok {
				fmt.Println(msg)
			} else {
				fmt.Println("channel is closed")
			}

		case msg, ok := <-channelB:
			if ok {
				fmt.Println(msg)
			} else {
				fmt.Println("channel is closed")
			}

		case msg, ok := <-channelC:
			if ok {
				fmt.Println(msg)
			} else {
				fmt.Println("channel is closed")
			}
		}
	}

	wg.Wait()
	close(channelA)
	close(channelB)
	close(channelC)

}
