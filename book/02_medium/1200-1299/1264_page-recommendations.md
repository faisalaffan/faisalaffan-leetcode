# 1264 — Page Recommendations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func pageRecommendations(userID int, likes []like, friendships []friend) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1264: Page Recommendations
// https://leetcode.com/problems/page-recommendations/
// Difficulty: Medium [Paid]

// Recommend pages liked by friends of user1 but not liked by user1.

// Time: O(n log n)
// Space: O(n)

type like struct {
	userID int
	pageID int
}

type friend struct {
	user1 int
	user2 int
}

func pageRecommendations(userID int, likes []like, friendships []friend) []int {
  // HashMap: O(1) lookup
	likedByUser := make(map[int]bool)
	for _, l := range likes {
		if l.userID == userID {
			likedByUser[l.pageID] = true
		}
	}

  // HashMap: O(1) lookup
	friends := make(map[int]bool)
	for _, f := range friendships {
		if f.user1 == userID {
			friends[f.user2] = true
		}
		if f.user2 == userID {
			friends[f.user1] = true
		}
	}

  // HashMap: O(1) lookup
	recommend := make(map[int]bool)
	for _, l := range likes {
		if friends[l.userID] && !likedByUser[l.pageID] {
			recommend[l.pageID] = true
		}
	}

  // Alokasi slice
	result := make([]int, 0, len(recommend))
	for p := range recommend {
		result = append(result, p)
	}
  // Sort O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	likes := []like{
		{1, 101}, {1, 102}, {2, 101}, {2, 103}, {3, 102},
	}
	friendships := []friend{
		{1, 2}, {1, 3},
	}
	fmt.Printf("%v (expected: [103] or [102 103])\n",
		pageRecommendations(1, likes, friendships))

	fmt.Printf("%v (expected: [])\n",
		pageRecommendations(2, likes, friendships))
}
```
