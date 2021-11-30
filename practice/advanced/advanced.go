package main

import (
	"fmt"
	"os"
	"time"
)

func go_route_f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func chan_sync_worker(done chan bool) {
	fmt.Println("working....")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pong chan<- string) {
	msg := <-pings
	pong <- msg
}

func worker_groups(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func wait_groups(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	//This is where the sorting logic is occuring
	return len(s[i]) < len(s[j])
	//return s[i] < s[j]
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func mayPanic() {
	panic("a problem")
}

func main() {

	////////////////////////////////
	// goroutines
	////////////////////////////////
	/*
		go_route_f("direct")

		go go_route_f("goroutine")

		go func(msg string) {
			fmt.Println(msg)
		}("going")

		time.Sleep(time.Second) // When running go routines, make sure to sleep so the the async processes have time to finish
		fmt.Println("done")
	*/

	////////////////////////////////////////////////////////////////
	// Channels
	////////////////////////////////////////////////////////////////
	/*
		messages := make(chan string)

		go func() { messages <- "ping" }()
		fmt.Println("messages address:", messages)
		//fmt.Println("messages value:", <-messages)
		// you can only send the message from the channel once

		msg := <-messages
		fmt.Println("msg value:", msg)

		//messg := make(chan string)
		// messg <- "test"
		//fmt.Println("msssg value:", messg)
	*/

	////////////////////////////////
	// Channel Buffering
	////////////////////////////////
	/*
		messages := make(chan string, 2)

		messages <- "First"
		messages <- "Second"
		// By default channels are unbuffered, the require a reciever from a goroutine.
		// Create a buffered channel to recieve meesages without a goroutine.

		fmt.Println("First: ", <-messages)
		fmt.Println("First: ", <-messages)
	*/

	////////////////////////////////
	// Channel Syncronization
	////////////////////////////////
	/*
		done := make(chan bool, 1)
		go chan_sync_worker(done)

		<-done // without this notification from the channel the code would end before it started
	*/

	////////////////
	// Channel direction
	////////////////

	/*
		pings := make(chan string, 1)
		pongs := make(chan string, 1)
		ping(pings, "Passed Message")
		pong(pings, pongs)

		fmt.Println(<-pongs)
	*/

	////////////////////////////////////////////////////////////////
	// Select
	////////////////////////////////////////////////////////////////
	/*

		c1 := make(chan string, 1)
		c2 := make(chan string, 1)

		go func() {
			time.Sleep(1 * time.Second)
			c1 <- "message 1"
		}()

		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "message 2"
		}()

		for i := 0; i < 2; i++ {
			select {
			case res1 := <-c1:
				fmt.Println(res1)
			case res2 := <-c2:
				fmt.Println(res2)
			}
		}
	*/
	////////////////////////////////////////////////
	//// Timeouts
	////////////////////////////////////////////////
	/*
		c1 := make(chan string, 1)

		go func() {
			time.Sleep(2 * time.Second)
			c1 <- "result 1"
		}()

		select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1")
		}

		c2 := make(chan string, 1)

		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "result 2"
		}()

		select {
		case res := <-c2:
			fmt.Println(res)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 2")

		}
	*/

	////////////////////////////////
	// nonblocking channel operations
	////////////////////////////////
	/*
		messages := make(chan string)
		signals := make(chan string)

		select {
		case msg := <-messages:
			fmt.Println("recieved message:", msg)
		default:
			fmt.Println("no message recieved")
		}

		msg := "hi"
		select {
		case messages <- msg:
			fmt.Println("sent message:", msg)
		default:
			fmt.Println("no message sent")
		}

		select {
		case msg := <-messages:
			fmt.Println("recieved message:", msg)
		case sig := <-signals:
			fmt.Println("recieved signal:", sig)
		default:
			fmt.Println("no activity")

		}
	*/
	////////////////////////////////////////////////////////////////
	// Closing Channels
	////////////////////////////////////////////////////////////////
	/*
		jobs := make(chan int, 5)
		done := make(chan bool)

		go func() {
			for {
				i, more := <-jobs
				if more {
					fmt.Println("recieved job: ", i)
				} else {
					fmt.Println("Recieved all jobs")
					done <- true
					return
				}
			}
		}()

		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("Sent Job:", j)
			time.Sleep(2 * time.Second)
		}
		close(jobs)
		fmt.Println("Closed all jobs")

		//<-done Done isnt necessary because of the pause..... interesting
	*/

	////////////////////////////////////////////////////////////////
	// range over Channel
	////////////////////////////////////////////////////////////////
	/*

		queue := make(chan string, 2)
		queue <- "one"
		queue <- "two"
		close(queue)

		for elem := range queue {
			fmt.Println(elem)
		}
	*/
	// You can still recieve when a queue is closed

	////////////////////////////////////////////////////////////////
	// Timers
	////////////////////////////////////////////////////////////////
	/*
		timer1 := time.NewTimer(2 * time.Second)

		fmt.Println("Timer 1 Fired", <-timer1.C)

		timer2 := time.NewTimer(time.Second)

		go func() {
			fmt.Println("Timer 2 Fired", <-timer2.C)
		}()
		stop2 := timer2.Stop()
		if stop2 {
			fmt.Println("Timer 2 Stopped")
		}

		time.Sleep(2 * time.Second)
	*/

	////////////////////////////////
	// Ticker
	////////////////////////////////

	/*
		ticker := time.NewTicker(500 * time.Millisecond)
		done := make(chan bool)

		go func() {
			for {
				select {
				case <-done:
					return
				case t := <-ticker.C:
					fmt.Println("Ticker at:", t)
				}
			}
		}()

		time.Sleep(1600 * time.Millisecond)
		ticker.Stop()
		done <- true
		fmt.Println("Ticker stopped")
	*/

	////////////////////////////////
	// Worker Pool
	////////////////////////////////
	/*
		const num_jobs = 5
		jobs := make(chan int, num_jobs)
		results := make(chan int, num_jobs)

		for w := 1; w <= 3; w++ {
			go worker_groups(w, jobs, results)
		}

		for i := 0; i <= num_jobs; i++ {
			jobs <- i
		}
		close(jobs)

		for a := 1; a <= num_jobs; a++ {
			fmt.Println("Return", <-results)
		}
	*/
	////////////////////////////////
	// WaitGroups
	////////////////////////////////
	/*
		var wg sync.WaitGroup

		for i := 1; i <= 5; i++ {
			wg.Add(1)

			i := i

			go func() {
				defer wg.Done()
				wait_groups(i)
			}()
		}

		wg.Wait()
	*/

	////////////////////////////////
	// Rate Limiting
	////////////////////////////////
	/*

		request := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			request <- i
		}

		close(request)

		limiter := time.Tick(200 * time.Millisecond)

		for req := range request {
			<-limiter
			fmt.Println("request", req, time.Now())
		}

		bursty_limit := make(chan time.Time, 3)

		for i := 0; i < 3; i++ {
			bursty_limit <- time.Now()
		}

		go func() {
			for t := range time.Tick(200 * time.Millisecond) {
				bursty_limit <- t
			}
		}()

		bursty_request := make(chan int, 5)
		// Cant add 0 to a channel
		for i := 1; i <= 5; i++ {
			bursty_request <- i
		}
		close(bursty_request)

		for r := range bursty_request {
			<-bursty_limit
			fmt.Println("request", r, time.Now())
		}
	*/

	////////////////////////////////
	// Atomic Counter
	////////////////////////////////
	/*

		var ops uint64
		var wg sync.WaitGroup

		for i := 0; i < 50; i++ {
			wg.Add(1)

			go func() {
				for c := 0; c < 1000; c++ {
					atomic.AddUint64(&ops, 1)
					//ops += 1 This will not give you the answer you are looking for
				}
				wg.Done()
			}()
		}

		wg.Wait()
		fmt.Println("ops:", ops)
	*/

	////////////////////////////////
	// Mutex
	////////////////////////////////
	/*
		var state = make(map[int]int)
		var mutex = &sync.Mutex{}

		var readOps uint64
		var writeOps uint64

		for r := 0; r < 100; r++ {
			go func() {
				// fmt.Println("Starting Read")
				total := 0
				for {
					key := rand.Intn(5)
					mutex.Lock()
					total += state[key]
					mutex.Unlock()
					atomic.AddUint64(&readOps, 1)
					time.Sleep(time.Millisecond)
				}
			}()
		}

		for w := 0; w < 10; w++ {
			go func() {
				//fmt.Println("Starting Write")
				for {
					key := rand.Intn(5)
					val := rand.Intn(100)
					mutex.Lock()
					state[key] = val
					mutex.Unlock()
					atomic.AddUint64(&writeOps, 1)
					time.Sleep(time.Millisecond)
				}

			}()
		}

		time.Sleep(time.Second)

		readOpsFinal := atomic.LoadUint64(&readOps)
		fmt.Println("Read Ops:", readOpsFinal)
		writeOpsFinal := atomic.LoadUint64(&writeOps)
		fmt.Println("Write Ops:", writeOpsFinal)

		mutex.Lock()
		//state can still be accessed without a lock
		fmt.Println("State:", state)
		// Program will still compile even if you dont release the lock at the end
		mutex.Unlock()
	*/
	////////////////////////////////////////////////////////////////
	// Stateful Goroutine
	////////////////////////////////////////////////////////////////
	/*
		var readOps uint64
		var writeOps uint64

		reads := make(chan readOp)
		writes := make(chan writeOp)

		go func() {
			var state = make(map[int]int)
			for {
				select {
				case read := <-reads:
					read.resp <- state[read.key]
				case write := <-writes:
					state[write.key] = write.val
					write.resp <- true
				}
			}
		}()

		for r := 0; r < 100; r++ {
			go func() {
				for {
					read := readOp{
						key:  rand.Intn(5),
						resp: make(chan int)}
					reads <- read
					<-read.resp
					atomic.AddUint64(&readOps, 1)
					time.Sleep(time.Millisecond)
				}
			}()
		}

		for w := 0; w < 10; w++ {
			go func() {
				for {
					write := writeOp{
						key:  rand.Intn(5),
						val:  rand.Intn(100),
						resp: make(chan bool)}
					writes <- write
					<-write.resp
					atomic.AddUint64(&writeOps, 1)
					time.Sleep(time.Millisecond)
				}
			}()
		}

		time.Sleep(time.Second)
		readOpsFinal := atomic.LoadUint64(&readOps)
		fmt.Println("readOpsFinal: ", readOpsFinal)
		writeOpsFinal := atomic.LoadUint64(&writeOps)
		fmt.Println("writeOpsFinal: ", writeOpsFinal)
	*/

	////////////////////////////////////////////////////////////////
	// Sort
	////////////////////////////////////////////////////////////////
	/*
		strs := []string{"c", "a", "b"}
		fmt.Println("isSorted:", sort.StringsAreSorted(strs))
		sort.Strings(strs)
		fmt.Println("strs:", strs)

		ints := []int{4, 2, 67}
		fmt.Println("isSorted:", sort.IntsAreSorted(ints))
		sort.Ints(ints)
		fmt.Println("ints:", ints)

		isSorted := sort.IntsAreSorted(ints)
		fmt.Println("isSorted:", isSorted)
	*/

	////////////////////////////////
	// Sorting by Function
	////////////////////////////////
	/*
		fruits := []string{"peach", "banana", "kiwi"}
		sort.Sort(byLength(fruits))
		fmt.Println(fruits)
	*/

	///////////////////////////////
	// Panic
	///////////////////////////////
	/*
		fmt.Println("Going to Panic")
		//panic("a problem")

		_, err := os.Create("/temp/file")
		if err != nil {
			panic(err)
		}
	*/

	////////////////////////////////
	// Defer
	////////////////////////////////
	/*
		f := createFile("/tmp/defer.txt")
		defer closeFile(f)
		writeFile(f)
	*/

	////////////////////////////////
	// Recover
	////////////////////////////////

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")

}
