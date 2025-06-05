package main

import (
	"fmt"
	"sync"
	"time"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func concurrencyDemo() {
	done := make(chan string)
	go func() {
		time.Sleep(time.Second)
		done <- "goroutine finished"
	}()
	fmt.Println(<-done)
}

// mapDemo demonstrates Go's map behaviors. It highlights
// initialization, read/write semantics, and iteration order.
// Each code block is annotated with its experimental intent.
func mapDemo() {
	fmt.Println("\nMap basics:")

	// 1. Start with a nil map. It's read-only and can't be written to.
	var scores map[string]int // nil map
	fmt.Printf("nil map? %v\n", scores == nil)
	// scores["alice"] = 10 // would panic: assignment to entry in nil map

	// 2. Create a map using make so we can write to it.
	scores = make(map[string]int)
	scores["alice"] = 10
	scores["bob"] = 20

	// 3. Accessing a key returns the zero value when absent.
	v, ok := scores["carol"] // carol is not present
	fmt.Printf("carol => %d (found? %v)\n", v, ok)

	// 4. Delete a key and show resulting length.
	delete(scores, "alice")
	fmt.Printf("after delete len=%d\n", len(scores))

	// 5. Map iteration order is not guaranteed. Run twice to show.
	fmt.Println("iteration order example:")
	for i := 0; i < 2; i++ {
		for name, score := range scores {
			fmt.Printf("%s=%d ", name, score)
		}
		fmt.Println()
	}

	// 6. Map literals for quick initialization.
	preset := map[string]int{"dan": 40, "emma": 90}
	fmt.Println("literal map:", preset)
}

// structDemo showcases struct usage including zero values, literals,
// pointer modification, and anonymous structs. Each block is annotated
// to explain the experimental goal.
type person struct {
	name string
	age  int
}

func structDemo() {
	fmt.Println("\nStruct basics:")

	// 1. Zero value of a struct. Fields take their zero values.
	var p person
	fmt.Printf("zero value: %#v\n", p)

	// 2. Initialize using a struct literal.
	p2 := person{name: "Alice", age: 30}
	fmt.Printf("literal: %#v\n", p2)

	// 3. Modify a struct via pointer to demonstrate reference semantics.
	ptr := &p2
	ptr.age++
	fmt.Printf("after pointer update: %#v\n", p2)

	// 4. Anonymous struct for ad-hoc grouping of values.
	anon := struct {
		x int
		y int
	}{x: 1, y: 2}
	fmt.Printf("anonymous: %#v\n", anon)
}

// varDemo illustrates several forms of variable declaration
// and recommends when to use each style.
func varDemo() {
	fmt.Println("\nVariable declarations:")

	// 1. Declare with var and explicit type for zero value variables
	//    when a variable will be assigned later.
	var a int
	a = 1
	fmt.Println("var with type ->", a)

	// 2. Use short declaration := inside functions for most cases
	//    where the type can be inferred.
	b := "quick"
	fmt.Println(":= inferred ->", b)

	// 3. Declare and initialize with var when explicit type improves clarity
	var c float64 = 3.14
	fmt.Println("var with init ->", c)

	// 4. Group related declarations in a block when many variables
	//    share the same context.
	var (
		d int
		e string = "grouped"
	)
	fmt.Println("block ->", d, e)
}

// interfaceDemo shows how the empty interface can hold any type and
// why type assertions are needed. This pattern is generally used when
// the value's concrete type is unknown in advance.
func interfaceDemo() {
	fmt.Println("\ninterface{} usage:")

	var any interface{}

	// Store different types. The actual type is erased until asserted.
	any = 42
	fmt.Printf("value: %v type: %T\n", any, any)

	any = "hello"
	fmt.Printf("value: %v type: %T\n", any, any)

	// Use a type assertion to retrieve the underlying value safely.
	if s, ok := any.(string); ok {
		fmt.Println("asserted string length ->", len(s))
	}

	// Type switches are convenient for handling several possibilities.
	any = 99.9
	switch v := any.(type) {
	case int:
		fmt.Println("got int", v)
	case float64:
		fmt.Println("got float64", v)
	default:
		fmt.Println("unknown type")
	}
}

// channelLockDemo demonstrates using a channel as a binary semaphore
// to control access to a critical section. It's a lightweight
// alternative to sync.Mutex but should be released quickly.
func channelLockDemo() {
	fmt.Println("\nChannel as lock:")

	lock := make(chan struct{}, 1) // capacity 1 for mutual exclusion

	critical := func(id int) {
		lock <- struct{}{} // acquire
		fmt.Printf("worker %d in critical section\n", id)
		time.Sleep(10 * time.Millisecond)
		<-lock // release
	}

	for i := 0; i < 2; i++ {
		go critical(i)
	}
	time.Sleep(50 * time.Millisecond)
}

// channelMessageDemo shows typical message passing via channels,
// including buffered channels and closure behavior.
func channelMessageDemo() {
	fmt.Println("\nChannel messaging:")

	// Buffered channel allows sending without immediate receiver.
	ch := make(chan string, 2)
	ch <- "a"
	ch <- "b"
	close(ch) // signal no more values

	for msg := range ch {
		fmt.Println("received", msg)
	}
}

// goroutineDemo highlights best practices: use WaitGroup to wait for
// goroutines to finish and capture loop variables correctly.
func goroutineDemo() {
	fmt.Println("\nGoroutine best practices:")

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		i := i // capture loop variable
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("goroutine %d working\n", i)
			time.Sleep(20 * time.Millisecond)
		}()
	}
	wg.Wait()
}

func main() {
	fmt.Println("## Basic Go Demo ##")
	fmt.Println("Hello, Go!")

	fmt.Println("\nFactorial example:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}

	fmt.Println("\nConcurrency example:")
	concurrencyDemo()

	fmt.Println("\nVariable example:")
	varDemo()

	fmt.Println("\ninterface{} example:")
	interfaceDemo()

	fmt.Println("\nChannel lock example:")
	channelLockDemo()

	fmt.Println("\nChannel messaging example:")
	channelMessageDemo()

	fmt.Println("\nGoroutine example:")
	goroutineDemo()

	fmt.Println("\nMap example:")
	mapDemo()

	fmt.Println("\nStruct example:")
	structDemo()
}
