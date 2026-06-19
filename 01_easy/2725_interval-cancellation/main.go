package main

// LeetCode #2725: Interval Cancellation
// https://leetcode.com/problems/interval-cancellation/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript setInterval problem, adapted to Go.

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	cancel := IntervalCancellation(func() {
		count++
		fmt.Println("tick", count)
	}, 50)
	time.Sleep(200 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("total ticks:", count)
}

func IntervalCancellation(fn func(), delay int) func() {
	ticker := time.NewTicker(time.Duration(delay) * time.Millisecond)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fn()
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	return func() {
		close(done)
	}
}
