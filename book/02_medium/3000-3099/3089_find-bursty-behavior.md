# 3089 — Find Bursty Behavior

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func findBurstyBehavior(posts [][]int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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

  // HashMap: O(1) lookup
	userPosts := make(map[int][]int)
	for _, p := range posts {
		userID, timestamp := p[0], p[1]
		userPosts[userID] = append(userPosts[userID], timestamp)
	}

	var bursty []int
	for uid, timestamps := range userPosts {
  // Sort O(n log n)
		sort.Ints(timestamps)
		for i := k - 1; i < len(timestamps); i++ {
			if timestamps[i]-timestamps[i-k+1] <= 100 {
				bursty = append(bursty, uid)
				break
			}
		}
	}

  // Sort O(n log n)
	sort.Ints(bursty)
	return bursty
}

func main() {
	fmt.Println(findBurstyBehavior([][]int{{1, 10}, {1, 20}, {1, 30}, {2, 5}, {2, 200}}, 3)) // Expected: [1]
	fmt.Println(findBurstyBehavior([][]int{{1, 1}, {2, 2}, {3, 3}}, 2))                       // Expected: []
}
```
