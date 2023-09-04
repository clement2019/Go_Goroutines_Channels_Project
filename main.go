package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//Now lets create the function welcome()

func welcome(s string, times int) {

	for i := 0; i < times; i++ {

		//inject the time lag a 100ms delay
		time.Sleep(time.Millisecond * 100)
		fmt.Println(i, s)
	}

}

// now lets do another transation casling functions for credit thats adds 11 to the balance 5 times
var balance int64

func credits(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()

	for i := 0; i < 10; i++ {
		// Add 100 to balance atomically
		atomic.AddInt64(&balance, 100)

		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After crediting, balance is",
			balance)

	}
}

func debits(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// deducts -100 from balance atomically
		atomic.AddInt64(&balance, -100)
		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After debiting, balance is",
			balance)

	}

}

//lets now introduced Channels

// ---send data into a channel---
func sendData(ch chan string) {
	fmt.Println("Sending a string into channel...")
	// comment out the following line // time.Sleep(2 * time.Second)
	ch <- "Hello"
	fmt.Println("String has been retrieved from channel...")
}

// ---getting data from the channel---
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("String retrieved from channel:", <-ch)
}

///further examples of how goroutine,channels works toegther.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}


//iterating through the channels


	

//Now call the main function while calling the welcome function defined above to run two set of methods concurrently

func main() {

//iterating throug  channels



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

//////another one starts

	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	////Another one starts
	var wg sync.WaitGroup
	balance = 200
	fmt.Println("Initial balance is", balance)

	wg.Add(1) // add 1 to the WaitGroup counter
	go credits(&wg)

	wg.Add(1) // add 1 to the WaitGroup counter
	go debits(&wg)

	wg.Wait() // blocks until WaitGroup counter is 0

	fmt.Println("Final balance is", balance)
	fmt.Scanln()

	//////another one starts

	go welcome("Hello guys", 3)
	go welcome("Welcome back", 2)
	//fmt.Scanln()

}
