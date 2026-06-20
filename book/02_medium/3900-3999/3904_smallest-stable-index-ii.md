# 3904 — Smallest Stable Index Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestStableIndexIi(nums []int, k int) int
```

> **💡 Hint:** Precompute suffix minimum, iterate prefix maximum.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3904: Smallest Stable Index II
// https://leetcode.com/problems/smallest-stable-index-ii/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Precompute suffix minimum, iterate prefix maximum.
// Check if max(nums[0..i]) - min(nums[i..n-1]) <= k.

import "fmt"

func SmallestStableIndexIi(nums []int, k int) int {
	n := len(nums)
  // Alokasi slice integer
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(nums[i], suffixMin[i+1])
	}

	prefixMax := 0
	for i := 0; i < n; i++ {
		if nums[i] > prefixMax {
			prefixMax = nums[i]
		}
		if prefixMax-suffixMin[i] <= k {
			return i
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println(SmallestStableIndexIi([]int{5, 0, 1, 4}, 3)) // Expected: 3

	// Example 2
	fmt.Println(SmallestStableIndexIi([]int{3, 2, 1}, 1)) // Expected: -1

	// Example 3
	fmt.Println(SmallestStableIndexIi([]int{0}, 0)) // Expected: 0
}
```
