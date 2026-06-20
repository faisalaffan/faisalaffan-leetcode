# 2948 — Make Lexicographically Smallest Array By Swapping Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func lexicographicallySmallestArray(nums []int, limit int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2948: Make Lexicographically Smallest Array by Swapping Elements
// https://leetcode.com/problems/make-lexicographically-smallest-array-by-swapping-elements/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lexicographicallySmallestArray([]int{1, 5, 3, 9, 8}, 2))
	fmt.Println(lexicographicallySmallestArray([]int{1, 7, 6, 18, 2, 1}, 3))
	fmt.Println(lexicographicallySmallestArray([]int{1, 2, 3}, 0))
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
  // Alokasi slice integer
	idx := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range idx {
		idx[i] = i
	}
  // Custom sort dengan comparator
	sort.Slice(idx, func(i, j int) bool {
		return nums[idx[i]] < nums[idx[j]]
	})
  // Alokasi slice integer
	ans := make([]int, n)
	for i := 0; i < n; {
		j := i + 1
		for j < n && nums[idx[j]]-nums[idx[j-1]] <= limit {
			j++
		}
  // Alokasi slice integer
		t := make([]int, j-i)
		copy(t, idx[i:j])
  // Urutkan secara ascending — O(n log n)
		sort.Ints(t)
		for k := i; k < j; k++ {
			ans[t[k-i]] = nums[idx[k]]
		}
		i = j
	}
	return ans
}
```
