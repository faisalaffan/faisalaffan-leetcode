package main

// LeetCode #2621: Sleep
// https://leetcode.com/problems/sleep/
// Difficulty: Easy
// Time: O(millis) | Space: O(1)
// Note: JavaScript async problem, adapted to Go.

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sleep(100)
	fmt.Println("Slept for", time.Since(start).Milliseconds(), "ms")
}

func sleep(millis int) {
	time.Sleep(time.Duration(millis) * time.Millisecond)
}
