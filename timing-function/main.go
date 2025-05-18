package main

import (
	"fmt"
	"time"
)

func withTiming(f func()) func() {
	return func() {
		start := time.Now()
		f()
		fmt.Printf("Execution time: %s\n", time.Since(start))

	}
}

func doWork() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Work done")
}

func main() {
	timed := withTiming(doWork)
	timed()
}