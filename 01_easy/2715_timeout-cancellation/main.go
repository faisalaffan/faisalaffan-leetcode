package main

// LeetCode #2715: Timeout Cancellation
// https://leetcode.com/problems/timeout-cancellation/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns a cancel function.

import (
	"fmt"
	"time"
)

func main() {
	cancel := TimeoutCancellation(func() { fmt.Println("executed") }, 100)
	time.Sleep(50 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done")
}

func TimeoutCancellation(fn func(), delay int) func() {
	timer := time.AfterFunc(time.Duration(delay)*time.Millisecond, fn)
	return func() {
		timer.Stop()
	}
}
