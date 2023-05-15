package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendData(ch chan string) {
	fmt.Println("Sending a string into channel...")
	// time.Sleep(2 * time.Second)
	ch <- "Hello"
	fmt.Println("String has been retrieved from channel...")
}

func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("String retreived from channel:", <-ch)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	fmt.Println("Done and can continue to do other work")
}

func basicExample() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	fmt.Scanln()
}

func channelCommunication() {
	s := []int{}
	sliceSize := 10
	for i := 0; i < sliceSize; i++ {
		s = append(s, rand.Intn(100))
	}

	c := make(chan int)
	partSize := 2
	parts := sliceSize / partSize
	i := 0
	for i < parts {
		go sum(s[i*partSize:(i+1)*partSize], c)
		i += 1
	}

	i = 0
	total := 0
	for i < parts {
		partialSum := <-c
		fmt.Println("Partial Sum: ", partialSum)
		total += partialSum
		i += 1
	}
	fmt.Println("Total: ", total)
	fmt.Scanln()
}

func fib(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
		time.Sleep(2 * time.Second)
	}
	close(c)
}

func counter(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func asyncWait() {
	c1 := make(chan int)
	c2 := make(chan int)

	go fib(10, c1)
	go counter(10, c2)

	c1Closed := false
	c2Closed := false

	go func() {
		for {
			select {
			case n, ok := <-c1:
				if !ok {
					// channel closed and drained
					c1Closed = true
					if c1Closed && c2Closed {
						return
					}
				} else {
					fmt.Println("fib()", n)
				}
			case m, ok := <-c2:
				if !ok {
					// channel closed and drained
					c2Closed = true
					if c1Closed && c2Closed {
						return
					}
				} else {
					fmt.Println("counter()", m)
				}
			}
		}
	}()

	fmt.Println("Continue")
	fmt.Scanln()
}

func bufferedChannel() {

	s := []int{}
	sliceSize := 10
	for i := 0; i < sliceSize; i++ {
		s = append(s, rand.Intn(100))
	}

	c := make(chan int, 5)
	partSize := 2
	parts := sliceSize / partSize
	i := 0
	for i < parts {
		go sum(s[i*partSize:(i+1)*partSize], c)
		i += 1
	}

	i = 0
	total := 0
	for i < parts {
		partialSum := <-c
		fmt.Println("Partial Sum: ", partialSum)
		total += partialSum
		i += 1
	}
	fmt.Println("Total: ", total)
	fmt.Scanln()
}

func main() {
	basicExample()
	// channelCommunication()
	// asyncWait()
	// bufferedChannel()
}
