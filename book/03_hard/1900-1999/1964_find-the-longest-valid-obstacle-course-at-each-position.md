# 1964 — Find The Longest Valid Obstacle Course At Each Position

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestObstacleCourseAtEachPosition(obstacles []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1964: Find the Longest Valid Obstacle Course at Each Position
// https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/
// Difficulty: Hard
// LIS variant with non-decreasing constraint.
// Patience sorting: binary search for first element > h, replace with h.

import (
	"fmt"
	"sort"
)

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
  // Alokasi slice integer
	ans := make([]int, n)
  // Alokasi slice integer
	tails := make([]int, 0, n)

	for i, h := range obstacles {
		// Find first element in tails > h (strictly greater)
		// sort.Search finds first index where predicate is true
		idx := sort.Search(len(tails), func(k int) bool { return tails[k] > h })
		if idx == len(tails) {
			tails = append(tails, h)
		} else {
			tails[idx] = h
		}
		ans[i] = idx + 1
	}
	return ans
}

func main() {
	fmt.Println(longestObstacleCourseAtEachPosition([]int{1, 2, 3, 2}))          // Expected: [1 2 3 3]
	fmt.Println(longestObstacleCourseAtEachPosition([]int{2, 2, 1}))             // Expected: [1 2 1]
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))   // Expected: [1 1 2 3 2 2]
}
```
