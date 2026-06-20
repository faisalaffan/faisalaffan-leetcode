# 3089 — Find Bursty Behavior

## Deskripsi

**Soal:** [3089. Find Bursty Behavior](https://leetcode.com/problems/find-bursty-behavior/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func findBurstyBehavior(posts [][]int, k int) []int`

## Solusi Go

```go
package main

// LeetCode #3089: Find Bursty Behavior
// https://leetcode.com/problems/find-bursty-behavior/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findBurstyBehavior(posts [][]int, k int) []int {
	n := len(posts)
  // Edge case: input kosong
	if n == 0 {
		return nil
	}

  // Membuat map untuk pencarian O(1): key → value
	userPosts := make(map[int][]int)
	for _, p := range posts {
		userID, timestamp := p[0], p[1]
		userPosts[userID] = append(userPosts[userID], timestamp)
	}

	var bursty []int
	for uid, timestamps := range userPosts {
		sort.Ints(timestamps)
		for i := k - 1; i < len(timestamps); i++ {
			if timestamps[i]-timestamps[i-k+1] <= 100 {
				bursty = append(bursty, uid)
				break
			}
		}
	}

	sort.Ints(bursty)
	return bursty
}

func main() {
	fmt.Println(findBurstyBehavior([][]int{{1, 10}, {1, 20}, {1, 30}, {2, 5}, {2, 200}}, 3)) // Expected: [1]
	fmt.Println(findBurstyBehavior([][]int{{1, 1}, {2, 2}, {3, 3}}, 2))                       // Expected: []
}
```
