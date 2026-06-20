# 1951 — All The Pairs With The Maximum Number Of Common Followers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaxCommonFollowers(relations [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n^2 * m) roughly, Space: O(n*m)  |  **Ruang:** O(n*m)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1951: All the Pairs With the Maximum Number of Common Followers
// https://leetcode.com/problems/all-the-pairs-with-the-maximum-number-of-common-followers/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// relations: [user_id, follower_id]
	relations := [][]int{{1, 3}, {2, 3}, {3, 4}, {1, 4}, {2, 4}, {1, 5}}
	fmt.Println(MaxCommonFollowers(relations))
}

// Time: O(n^2 * m) roughly, Space: O(n*m)
func MaxCommonFollowers(relations [][]int) [][]int {
  // HashMap: O(1) lookup
	followers := make(map[int]map[int]bool)
	for _, r := range relations {
		user, follower := r[0], r[1]
		if followers[user] == nil {
			followers[user] = make(map[int]bool)
		}
		followers[user][follower] = true
	}

  // Alokasi slice
	users := make([]int, 0, len(followers))
	for u := range followers {
		users = append(users, u)
	}
  // Sort O(n log n)
	sort.Ints(users)

	maxCommon := 0
  // Matriks 2D
	result := make([][]int, 0)

  // Linear scan O(n)
	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			a, b := users[i], users[j]
			common := 0
			for f := range followers[a] {
				if followers[b][f] {
					common++
				}
			}
			if common > maxCommon {
				maxCommon = common
				result = [][]int{{a, b}}
			} else if common == maxCommon && common > 0 {
				result = append(result, []int{a, b})
			}
		}
	}
	return result
}
```
