### Go Goroutines/Channels Project
This Go Goroutines / Channels Project although small in scope but it clearly shows how multitasking can be done concurrently in Go.
In this program, you have a function named welcome(), which takes in two arguments: a string to print on the console and the number of times the string must be printed. In the main() function, you call the welcome() function two times: first with the string "Hello guys", and again with the string "Welcome back".

## When you run the program, you see the following output:

0 Hello guys
1 Hello guys
2 Hello guys
0 Welcome back
1 Welcome back

The above outpouts is what you expected because the first call to the welcome() function must end before the second call can proceed. But what if you want to make the two calls run concurrently?.For any seasoned developer, you’re familiar with a concept known as threading, which allows you to implement concurrent operations — multiple functions can all be running at the same time. To call the welcome() function concurrently in Go, you use goroutines.

In Go, a goroutine is a lightweight thread managed by the Go runtime. To run a function as a goroutine, simply call it using the go keyword:

func main() {

	go welcome("Hello guys", 3)
	go welcome("Welcome back", 2)
    fmt.ScanIn()
}

The first statement calls the welcome() function as a goroutine. Essentially, it means “Run the welcome() function independently and immediately return control back to the calling statement.” The second statement does the same. Now you have two separate instances of the welcome() function running concurrently. The result may appear like this (you may get a different result):

0 Hello guys
0 Welcome back
1 Welcome back
1 Hello guys
2 Hello guys

//Now lets deal with shared resources

Although goroutines makes it really easy for you to perform concurrent opera- tions, you have to use them with care when your goroutines access shared resources (such as variables). For example, two Goroutines may concurrently access a single variable — one crediting some amount to it and another one debit- ing from it. It’s important to ensure that when one Goroutine is adding values to the variable, another one is prevented from accessing it until the other Goroutine is done with it.
In the following project I explain the problems posed by multiple Goroutines trying to access the same shared resource, as well as the various techniques that you can use to resolve them.

## Seeing how shared resources impact goroutines
Consider the following program where you have two functions:
» The credit() function adds 100 to balance a total of five times.
» The debit() function deducts 100 from balance a total of five times.
You called the credit() and debit() functions simultaneously as goroutines.


/ now lets do another transation casling functions for credit thats adds 11 to the balance 5 times
var balance int

func credits() {

	for i := 0; i < 5; i++ {
		balance += 100

		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After crediting, balance is",
			balance)
	}
}

func debits() {

	for i := 0; i < 5; i++ {
		balance -= 100

		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After debiting, balance is",
			balance)
	}

}

//Now call the main function while calling the welcome function defined above to run two set of methods concurrently

func main() {
	balance = 200
	fmt.Println("The initial balance", balance)
	go credits()
	go debits()
	fmt.Scanln()
	go welcome("Hello guys", 3)
	go welcome("Welcome back", 2)
	//fmt.Scanln()
	
}

## The output shows like what you see below:


The initial balance 200
After debiting, balance is 200
After debiting, balance is 100
After crediting, balance is 0
After crediting, balance is 100
After crediting, balance is 200
After debiting, balance is 300
After crediting, balance is 200
After debiting, balance is 300
After crediting, balance is 200
After debiting, balance is 200



Notice anything wrong? The most obvious error is in the second line. If you look at the first line, the initial balance is 200. On the second line, it says that, after crediting, the balance is still 200? It should be 300. Why is it wrong?
After crediting the balance by 100 (where the balance should now be updated to 300), there was a delay of a random amount of time before the balance could be printed on the console. In the midst of this delay, the debit() function had a chance to deduct 100 from the balance (where the balance now would become 200). When the turn comes for the credit() function to continue printing the balance,


## Accessing shared resources using mutual exclusion

As you can see from the previous section, when you have goroutines that access the same variable, it’s important to make sure that only one goroutine can access the variable at any one time. To do so, you need to use the concept of mutual exclu- sion (also commonly known as mutex).
To implement mutual exclusion, Go’s standard library provides the Mutex type (a struct) in the sync package. A Mutex is a mutual exclusion lock. A mutual exclusion lock is a technique to ensure exclusive access to shared data between threads of execution. When one Goroutine gains the mutual exclusion lock, other Goroutines need to wait until the lock is released.
The following program shows how you can use a Mutex object to enclose blocks of code to ensure that when one block is being executed by a goroutine, another block can’t execute it:

   as shown in the main.go file. omce you run the command go run main.go,
   
   
   var mutex = &sync.Mutex{}
func credits() {

	for i := 0; i < 5; i++ {
		mutex.Lock()
		balance += 100

		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After crediting, balance is",
			balance)
			mutex.Unlock()
	}
}

func debits() {

	for i := 0; i < 5; i++ {
		mutex.Lock()
		balance -= 100
	
		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After debiting, balance is",
			balance)
			mutex.Unlock()
	}

}
func main() {
	balance = 200
	fmt.Println("The initial balance", balance)
	go credits()
	go debits()
	fmt.Scanln()
    }
   
   
   
   
   
    the output is as shown below, once the Mutex for sysnchnormous transation has been called to lock and unlock transations

   The initial balance 200
After crediting, balance is 300
After crediting, balance is 400
After debiting, balance is 300
After debiting, balance is 200
After crediting, balance is 300
After crediting, balance is 400
After debiting, balance is 300
After debiting, balance is 200
After crediting, balance is 300
After debiting, balance is 200


## Using atomic counters for modifying shared resources
In addition to using the Mutex object for marking critical sections, you can per- form changes to shared variables in a thread-safe manner using atomic counters (routines that allow you to perform mathematical operations on variables one thread at a time). The atomic package provides low-level atomic memory primi- tives useful for implementing synchronization algorithms. So, instead of using the Mutex object to ensure that the balance variable is credited or debited cor- rectly, use the AddInt64() function, like the following statements in bold:

The AddInt64() function atomically adds a value to the specified variable (in which you pass the address of the variable to be modified using the & operator) and returns the new value. The AddInt64() function ensures that when one Goroutine is adding a value to the specified variable, no other Goroutine is allowed.
to modify the specified variable until the addition operation is done. It works just like the mutual exclusion lock example discussed in the previous section, without the need to call the Lock() and Unlock() functions of a Mutex object.
Note that the preceding code credits the balance variable 10 times as shown belowe
var balance int64

func credits() {

	for i := 0; i < 10; i++ {
		// Add 100 to balance atomically
		atomic.AddInt64(&balance,100)

		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After crediting, balance is",
			balance)
			
	}
}

func debits() {

	for i := 0; i < 5; i++ {
		// deducts -100 from balance atomically
		atomic.AddInt64(&balance,-100)
		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After debiting, balance is",
			balance)
			
	}

}

//Now call the main function while calling the welcome function defined above to run two set of methods concurrently

func main() {
	balance = 200
	fmt.Println("The initial balance", balance)
	go credits()
	go debits()
	fmt.Scanln()
	fmt.Println(balance)



The preceding program prints out the following output:

As shown with the output below

The initial balance 200
After crediting, balance is 200
After debiting, balance is 300
After debiting, balance is 200
After crediting, balance is 100
After debiting, balance is 200
After debiting, balance is 100
After crediting, balance is 0
After debiting, balance is 100
After crediting, balance is 100
After crediting, balance is 200
After crediting, balance is 300
After crediting, balance is 400
After crediting, balance is 500
After crediting, balance is 600
After crediting, balance is 700



## Synchronizing Goroutines and the application of WaitGroup

When you have multiple Goroutines all running at the same time, you need to be able to synchronize them so that you have a way to coordinate the execution of your code. For example, you may have two Goroutines fetching data from differ- ent web services, and you need to ensure that these two Goroutines finish execu- tion before you go ahead with the next block of code. In this scenario, there must be a way for you to know when the Goroutines are done.
If you check out the code snippets in the previous section, you see that there is a Scanln() function near the end of the main() function:

func main() {
    balance = 200
fmt.Println("Initial balance is", balance) go credit()
go debit()
fmt.Scanln()
    fmt.Println(balance)
}

If you don’t have the Scanln() function, the main() function will immediately exit after calling the two goroutines, and the balance printed will not be the correct final result.

The solution to this problem? Wait groups.
When you have multiple goroutines running and you want to know when they’ve completed, you can use a wait group. To use a wait group, you first create a variable of type sync.WaitGroup, like this:

  var wg sync.WaitGroup

Now when you call a goroutine, simply call the wg.Add() function to add 1 to the WaitGroup counter:

  wg.Add(1) // add 1 to the WaitGroup counter

When the goroutine that you’ve called has finished execution, call the wg.Done() function to decrement the WaitGroup counter by 1:

  wg.Done() // decrement 1 from the WaitGroup counter

When you want to wait for the completion of all goroutines, use the wg.Wait() function:

  wg.Wait() // blocks until WaitGroup counter is 0

The Wait() function blocks until the WaitGroup counter is 0. As soon as the Wait- Group counter reaches 0, this means that all goroutines have been completed and you can continue executing all ensuing statements.
Now that you understand how the WaitGroup works, let me show you how it can be applied to the earlier example:


package main
import ( "fmt"
"math/rand" "sync" "sync/atomic" "time"
)
var balance int64

func credits(wg *sync.WaitGroup) {
	// notify the WaitGroup when we are done 
	defer wg.Done()

	for i := 0; i < 10; i++ {
		// Add 100 to balance atomically
		atomic.AddInt64(&balance,100)

		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After crediting, balance is",
			balance)
			
	}
}

func debits(wg *sync.WaitGroup){
	// notify the WaitGroup when we are done 
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// deducts -100 from balance atomically
		atomic.AddInt64(&balance,-100)
		//inject the time lag
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("After debiting, balance is",
			balance)
			
	}

}

//Now call the main function while calling the welcome function defined above to run two set of methods concurrently

func main() {
var wg sync.WaitGroup
balance = 200
fmt.Println("Initial balance is", balance)

wg.Add(1) // add 1 to the WaitGroup counter 
go credits(&wg)

wg.Add(1) // add 1 to the WaitGroup counter 
go debits(&wg)

wg.Wait() // blocks until WaitGroup counter is 0
}


## In this example, I’m doing the following:
Creating a WaitGroup object
» Calling the wg.Add() function to increment the WaitGroup counter prior to
calling the credit() goroutine
» Calling the wg.Add() function to increment the WaitGroup counter prior to
calling the debit() goroutine

Calling the wg.Done() function in the credit() and debit() functions when I’m done either crediting or debiting the balance variable
» Calling the wg.Wait() function to wait for the completion of the two goroutines
When the two goroutines are finally done, the WaitGroup counter becomes zero, and all statements blocked on wg.Wait() are released. The program will now print out the balance:


when the program is run the following is the output

Initial balance is 200
After debiting, balance is 200
After crediting, balance is 100
After debiting, balance is 200
After debiting, balance is 100
After debiting, balance is 0
After crediting, balance is -100
After debiting, balance is 0
After crediting, balance is 0
After crediting, balance is 100
After crediting, balance is 200
After crediting, balance is 300
After crediting, balance is 400
After crediting, balance is 500
After crediting, balance is 600
After crediting, balance is 700
Final balance is 700



### Communication between Goroutines Using Channels

We have previousely dealt with Goroutines and its usage however, one of the key features of the Go programming language. Goroutines are executed independently of one another and are a great way to implement concurrent programming. However, very often, goroutines need a way to communicate with each other in order to work properly. In Go, you can get goroutines to communicate with one another through pipes known as channels. In this scenario  I show you how channel works.

## Understanding channels
In Go, channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values in another goroutine. Think of channels as temporary storage for passing values between goroutines.

To create a channel, you use the make() function, together with the chan keyword and the type of data that you want the channel to store for you, like this:

  ch := make(chan int)
To send a value into the channel, you use the <- operator. The following example writes the value 5 into the ch channel:

ch <- 5

To retrieve a value from the channel, simply put the channel variable to the right of the <- operator. The following example retrieves a value from the ch channel and assigns it to a variable named value:

value := <- ch

## How channels work
To understand how channels work, let’s start off with a very simple example. Consider the following program, where you have two functions 

— sendData() and getData():

package main
import (
    "fmt"
"time" 
)

//---send data into a channel---

func sendData(ch chan string) {
    fmt.Println("Sending a string into channel...")
    time.Sleep(2 * time.Second)
    ch <- "Hello"
}

//---getting data from the channel---

func getData(ch chan string) {
    fmt.Println("String retrieved from channel:", <-ch)
}
func main() {
    ch := make(chan string)

    go sendData(ch)
    go getData(ch)
    fmt.Scanln()
}


In the main() function, you first create a channel using the make() function, ch, with the type of channel specified (string). This means that the channel can only contain values of type string. You then call the sendData() and getData() func- tions as goroutines. In the sendData() function, you first print out the sentence Sending a string into channel. . . After a delay of two seconds, you insert a string into the channel using the <- operator.
If the channel variable is to the left of the <- operator, you’re sending a value into the channel.

 If the channel variable is to the right of the <- operator, you’re receiv- ing a value from the channel.
At the same time, when you run the sendData() function, you’re also running the getData() function. Here, with getData(), you’re trying to receive a value from the channel. Because there is currently no value in the channel (it won’t have any value in it until two seconds later), the getData() function will block. The moment a value is available in the channel, the getData() function will unblock and retrieve the value from the channel. Hence, the output of the program will look like this:


For this example, the type of channel you’re creating is known as an unbuffered channel. In an unbuffered channel, the sender blocks until the value has been received by a receiver.
When you try to retrieve a value from a channel and there is no value available, your code will block.
Let’s now make the following changes to the program (shown in bold):

Sending a string into channel...
[After a two-second delay]
String retrieved from channel: Hello


Notice that immediately after sending a value into the channel, the sendData() function is blocked. It will only resume after the value in the channel is retrieved by the getData() function.

When you send a value into an unbuffered channel, your code will block until the value is retrieved from the channel.

## How channels are used

When you know how a channel works, you probably want to see some of the practical uses of channels and why you would want to use them in the first place.
Suppose you have a function named sum() that sums up a slice of integer values:

func sum(s []int, c chan int) {
    sum := 0
for _, v := range s { sum += v
}
c <- sum }

The first argument to the sum() function is a slice of int values; the second argu- ment is a channel of type int. When the numbers in the array have been summed up, the sum is written to the channel.
To use the sum() function, let’s now generate ten random numbers and assign it to a variable named s:

Channels behave like queues: All items are retrieved in the same order that they were written (first in, first out). As each goroutine writes to the channel, it’s blocked until its values in the channels are received.
After the first goroutine has sent a value into the channel, it’s blocked until the value is retrieved from the channel. Meanwhile, the other four goroutines trying to also send in values to the channel will be blocked until the value is removed from the channel. They get the chance to send a value into the channel on a first- come, first-served basis.
Because you know that you have five separate goroutines (and, therefore, a total of five values to be written to the channel), you can write a loop and try to extract the values in the channel:

func main() {
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
	i += 1 }
	i= 0
	total := 0
	for i < parts {
			partialSum := <-c
			fmt.Println("Partial Sum: ", partialSum)
			total += partialSum
			i += 1
		}
		fmt.Println("Total: ", total)
        fmt.ScanIn()
        }


Each value in the channel represents the partial sum of the values in each array. It’s important to remember that when you send a value into an unbuffered chan- nel, the goroutine is blocked until the value is received by another function/goroutine. Likewise, when you’re reading a value from a channel, your code is blocked until the data is read from the channel. If the goroutines are taking a long time to sum up, the preceding code snippet will block until all the partial sums are retrieved.

This is the output below

Partial Sum:  142
Partial Sum:  85
Partial Sum:  34
Partial Sum:  38
Partial Sum:  117
Total:  416
Initial balance is 200



iterating through gthe channels
func fib(n int, c chan int) { 
	
	a, b := 1, 1
	for i := 0; i < n; i++ {
	c <- a // blocked until value is received from channel a, b = b, a + b
	time.Sleep(1 * time.Second)
	} 
	close(c)
}


func main() {
c := make(chan int) 

go fib(10, c)

for i := range c {    // read from channel until channel is closed 

fmt.Println(i) 
}
}




