# 0933 — Number Of Recent Calls

## Deskripsi

**Soal:** [0933. Number Of Recent Calls](https://leetcode.com/problems/number-of-recent-calls/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1) amortized. Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** Queue (antrian FIFO)

**Fungsi Solusi:** `func Constructor() RecentCounter`

## Solusi Go

```go
package main

// LeetCode #933: Number of Recent Calls
// https://leetcode.com/problems/number-of-recent-calls/
// Difficulty: Easy

import "fmt"

// RecentCounter counts recent requests within a 3000ms window.
type RecentCounter struct {
	queue []int
}

// Constructor creates a RecentCounter.
func Constructor() RecentCounter {
	return RecentCounter{queue: make([]int, 0)}
}

// Ping adds a new request and returns the count of requests in the last 3000ms.
// Time: O(1) amortized. Space: O(n).
func (rc *RecentCounter) Ping(t int) int {
	rc.queue = append(rc.queue, t)
	for rc.queue[0] < t-3000 {
		rc.queue = rc.queue[1:]
	}
	return len(rc.queue)
}

func main() {
	rc := Constructor()
	fmt.Println(rc.Ping(1))    // 1
	fmt.Println(rc.Ping(100))  // 2
	fmt.Println(rc.Ping(3001)) // 3
	fmt.Println(rc.Ping(3002)) // 3
}
```
