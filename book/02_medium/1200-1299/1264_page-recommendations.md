# 1264 — Page Recommendations

## Deskripsi

**Soal:** [1264. Page Recommendations](https://leetcode.com/problems/page-recommendations/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func pageRecommendations(userID int, likes []like, friendships []friend) []int`

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
	likedByUser := make(map[int]bool)
	for _, l := range likes {
		if l.userID == userID {
			likedByUser[l.pageID] = true
		}
	}

  // Membuat map untuk pencarian O(1): key → value
	friends := make(map[int]bool)
	for _, f := range friendships {
		if f.user1 == userID {
			friends[f.user2] = true
		}
		if f.user2 == userID {
			friends[f.user1] = true
		}
	}

  // Membuat map untuk pencarian O(1): key → value
	recommend := make(map[int]bool)
	for _, l := range likes {
		if friends[l.userID] && !likedByUser[l.pageID] {
			recommend[l.pageID] = true
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0, len(recommend))
	for p := range recommend {
		result = append(result, p)
	}
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
