# 2723 — Add Two Promises

## Deskripsi

**Soal:** [2723. Add Two Promises](https://leetcode.com/problems/add-two-promises/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #2723: Add Two Promises
// https://leetcode.com/problems/add-two-promises/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript Promise problem, adapted to Go. Adds two integer results asynchronously.

import (
	"fmt"
	"time"
)

func main() {
	p1 := func() int { time.Sleep(50 * time.Millisecond); return 10 }
	p2 := func() int { time.Sleep(100 * time.Millisecond); return 20 }
	result := AddTwoPromises(p1, p2)
	fmt.Println(result)
}

func AddTwoPromises(promise1 func() int, promise2 func() int) int {
	result1 := make(chan int)
	result2 := make(chan int)

	go func() { result1 <- promise1() }()
	go func() { result2 <- promise2() }()

	return <-result1 + <-result2
}
```
