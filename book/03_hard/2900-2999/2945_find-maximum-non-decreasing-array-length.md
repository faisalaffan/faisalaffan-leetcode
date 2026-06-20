# 2945 — Find Maximum Non Decreasing Array Length

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findMaximumLength(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Prefix Sum, Monotonic Stack

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2945: Find Maximum Non-decreasing Array Length
// https://leetcode.com/problems/find-maximum-non-decreasing-array-length/
// Difficulty: Hard
//
// DP + monotonic deque optimization.
// Partition array into contiguous groups, replace each group with its sum.
// Goal: resulting array is non-decreasing, maximize number of groups.
//
// Define:
//   f[i] = max groups for prefix ending at i-1 (i elements)
//   g[i] = minimum possible last group sum achieving f[i]
//   pref[i] = prefix sum of first i elements
//
// Transition: for j < i where sum(j..i-1) = pref[i]-pref[j] >= g[j],
//   f[i] = f[j] + 1, g[i] = pref[i]-pref[j]
//
// Optimization: deque maintains candidates sorted by g[j] and pref[j]+g[j].

import (
	"fmt"
)

func findMaximumLength(nums []int) int {
	n := len(nums)
  // Alokasi slice
	pref := make([]int64, n+1)
	for i, v := range nums {
		pref[i+1] = pref[i] + int64(v)
	}

  // Alokasi slice
	f := make([]int, n+1)
  // Alokasi slice
	g := make([]int64, n+1)
  // Alokasi slice
	deq := make([]int, 0, n+1)
	deq = append(deq, 0)

	head := 0
	for i := 1; i <= n; i++ {
		// Pop front: discard indices that are no longer optimal
		for head+1 < len(deq) && pref[i] >= pref[deq[head+1]]+g[deq[head+1]] {
			head++
		}

		j := deq[head]
		f[i] = f[j] + 1
		g[i] = pref[i] - pref[j]

		// Pop back: maintain monotonicity of pref[i]+g[i]
		for len(deq) > head && pref[i]+g[i] <= pref[deq[len(deq)-1]]+g[deq[len(deq)-1]] {
			deq = deq[:len(deq)-1]
		}
		deq = append(deq, i)
	}
	return f[n]
}

func main() {
	// Example: [2,3,1,4,5] -> 4
	// Partition: [2],[3],[1,4],[5] -> sums [2,3,5,5] (non-decreasing, length 4)
	fmt.Println(findMaximumLength([]int{2, 3, 1, 4, 5}))

	// Simple cases
	fmt.Println(findMaximumLength([]int{1, 2, 3}))
	fmt.Println(findMaximumLength([]int{5, 4, 3, 2, 1}))
	fmt.Println(findMaximumLength([]int{1, 1, 1}))
	fmt.Println(findMaximumLength([]int{1, 2, 1, 2}))
}
```
