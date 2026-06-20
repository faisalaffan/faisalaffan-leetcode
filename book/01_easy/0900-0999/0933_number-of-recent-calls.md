# 0933 — Number Of Recent Calls

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func Constructor() RecentCounter`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(1) amortized. Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

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
