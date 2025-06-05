package main

import (
	"fmt"
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

func main() {
	fmt.Println("## Basic Go Demo ##")
	fmt.Println("Hello, Go!")

	fmt.Println("\nFactorial example:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}

	fmt.Println("\nConcurrency example:")
	concurrencyDemo()

	fmt.Println("\nMap example:")
	mapDemo()
}
