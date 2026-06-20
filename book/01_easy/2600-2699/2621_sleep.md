# 2621 — Sleep

## Deskripsi

**Soal:** [2621. Sleep](https://leetcode.com/problems/sleep/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(millis)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
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
```
