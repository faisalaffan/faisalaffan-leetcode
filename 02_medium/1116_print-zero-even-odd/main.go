package main

// LeetCode #1116: Print Zero Even Odd
// https://leetcode.com/problems/print-zero-even-odd/
// Difficulty: Medium
//
// Approach: Three goroutines with channels for ordering
// Time: O(n)
// Space: O(1)

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	zeroCh := make(chan bool, 1)
	evenCh := make(chan bool, 1)
	oddCh := make(chan bool, 1)
	zeroCh <- true

	n := 5
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i++ {
			<-zeroCh
			fmt.Print(0)
			if i%2 == 0 {
				evenCh <- true
			} else {
				oddCh <- true
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-evenCh
			fmt.Print(i)
			zeroCh <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			<-oddCh
			fmt.Print(i)
			zeroCh <- true
		}
	}()

	wg.Wait()
	fmt.Println()
}
