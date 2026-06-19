package main

import (
	"fmt"
	"sync"
)

// LeetCode #1195: Fizz Buzz Multithreaded
// https://leetcode.com/problems/fizz-buzz-multithreaded/
// Difficulty: Medium

// Multithreaded FizzBuzz using goroutines and WaitGroup.
// Each goroutine prints one line at a time.

// Time: O(n)
// Space: O(n)

type FizzBuzz struct {
	n    int
	done chan struct{}
}

func NewFizzBuzz(n int) *FizzBuzz {
	return &FizzBuzz{n: n, done: make(chan struct{})}
}

func (fb *FizzBuzz) Start() []string {
	result := make([]string, fb.n)
	var wg sync.WaitGroup
	wg.Add(4)

	// fizzbuzz goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= fb.n; i++ {
			if i%15 == 0 {
				result[i-1] = "FizzBuzz"
			}
		}
	}()

	// fizz goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= fb.n; i++ {
			if i%3 == 0 && i%5 != 0 {
				result[i-1] = "Fizz"
			}
		}
	}()

	// buzz goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= fb.n; i++ {
			if i%5 == 0 && i%3 != 0 {
				result[i-1] = "Buzz"
			}
		}
	}()

	// number goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= fb.n; i++ {
			if i%3 != 0 && i%5 != 0 {
				result[i-1] = fmt.Sprintf("%d", i)
			}
		}
	}()

	wg.Wait()
	return result
}

func main() {
	fb := NewFizzBuzz(15)
	result := fb.Start()
	fmt.Printf("%v\n", result)
}
