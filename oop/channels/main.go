package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func basicChannel() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		defer wg.Done()
		i := <-ch // received data
		ch <- 945 // sending data
		fmt.Println(i)
	}()

	go func() {
		defer wg.Done()
		i := 234
		ch <- i // send data
		fmt.Println(<-ch)
	}()
	wg.Wait()
}

func senderReceiverUnbuffered() {
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) { // Receiver end
		defer wg.Done()
		i := <-ch // received data
		fmt.Println(i)
	}(ch)

	go func(ch chan<- int) { // sender end
		defer wg.Done()
		i := 234
		ch <- i // send data
	}(ch)
	wg.Wait()
}

func senderReceiverBuffered() {
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) { // receiver end
		defer wg.Done()
		// for i := range ch {
		// 	fmt.Println(i)
		// }
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
	}(ch)

	go func(ch chan<- int) {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	wg.Wait()
}

// it is not mendatory to have a reciver for a buffered channel
func bufferedWithNoReceiver() {
	ch := make(chan int, 1) // Buffered channel with capacity 1

    go func() {
        ch <- 1 // Goroutine চ্যানেলে ডেটা পাঠাবে
    }()
}

func portal1(ch1 chan string) {
	time.Sleep(2 * time.Second)
	ch1 <- "channel 1"
}

func portal2(ch2 chan string) {
	time.Sleep(1 * time.Second)
	ch2 <- "channel 2"
}

func selectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go portal1(ch1)
	go portal2(ch2)

	// select helps to see from what we are getting data
	select {
	case option1 := <-ch1:
		fmt.Println(option1)
	case option2 := <-ch2:
		fmt.Println(option2)
	} 
}



func main() {
	basicChannel()
	senderReceiverUnbuffered()
	senderReceiverBuffered()
	bufferedWithNoReceiver()
	selectStatement()
}
