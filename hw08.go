package main

import (
	"fmt"
	"runtime"
	"time"
)

func hw08() {
	fmt.Printf(
		"Goroutines: %d\n",
		runtime.NumGoroutine(),
	)

	// Channels to sync between goroutines
	var ch = make(chan struct{}) //created channel

	go func() { //call closure as goroutine
		fmt.Printf("Hello\n")
		ch <- struct{}{} // write to the channel (or here close(ch))
	}()

	<-ch //wait for read from the channel
	close(ch)

	//send data to many channels
	var start = make(chan struct{})
	for i := 0; i < 100; i++ {
		go func() {
			<-start
			// горутины не начнут работу,
			// пока не будут созданы все 10000
		}()
	}

	close(start) // only here all goroutins starts the jobs

	//timers
	timer := time.NewTimer(1 * time.Second)
	select {
	case data := <-ch:
		fmt.Printf("received: %v\n", data)
	case <-timer.C:
		fmt.Printf("failed to receive in 10s\n")
	}

	//tickers
	ticker := time.NewTicker(1*time.Second)
	for {
		select {
		case <- ticker.C:
			fmt.Println("do something")
			break
		}
		break
	}

	//we need write 01234, but the closure takes the value of i in the execution moment, so we will always have random string
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i)
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("")

	//try fix error in the prev sample, but with no luck, because goroutins executes in random order
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Print(j)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("")

	//fix that stuff with buffered channel, which stores 01234 values and each goroutine takes only one value and execute
	N := 5
	buffer := make(chan int, N)
	for i := 0; i < 5; i++ {
		buffer <- i
		go func() {
			i := <-buffer
			fmt.Print(i)
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("")
	close(buffer)
}

// errorString is a trivial implementation of error.
type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

func doParallelTasks(tasks []func() error, maxParallelCount int, maxErrors int) error {

	runChan := make(chan int, maxParallelCount)
	rsltChan := make(chan error)
	//executed := 0

	runJobs := func() {
		index := 0
		for {
			runChan <- index

			go func(index int) {
				fmt.Printf("Start task %v\n", index)
				rsltChan <- tasks[index]() // iterate over i < maxParallelCount, but use nextJobIndex() to fetch next job
				fmt.Printf("Finish task %v\n", index)
			} (index)

			index++

			if index == len(tasks)-1 {
				break
			}

			if maxErrors <= 0 {
				break
			}
		}
	}

	go runJobs() //run Goroutines one by one, we can control them from here via 'runChan' buffer

	executed := 0
	for {
		res := <-rsltChan
		<-runChan
		executed++

		if res != nil {
			maxErrors--
		}

		//fmt.Printf("Errors to exit: %v\n", maxErrors)
		if maxErrors <= 0 {
			return &ErrorString{"Max errors exceeded"}
		}

		//fmt.Printf("Executed: %v %v\n", executed, len(tasks))
		if executed >= len(tasks)-1 {
		    return nil
		}
	}

	return nil
}