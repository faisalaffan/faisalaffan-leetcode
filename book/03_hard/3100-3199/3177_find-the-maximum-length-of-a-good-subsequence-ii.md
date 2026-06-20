# 3177 — Find The Maximum Length Of A Good Subsequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumLength(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3177: Find the Maximum Length of a Good Subsequence II
// https://leetcode.com/problems/find-the-maximum-length-of-a-good-subsequence-ii/
// Difficulty: Hard
//
// A subsequence is "good" if at most k adjacent pairs have different values.
// Find the maximum possible length of a good subsequence.
//
// Approach: DP tracking best[val][k] and overall best[k].
//   best[val][k] = max length of good subsequence ending with value val
//                  using at most k diff-pairs.
//   global[k]     = max over all best[val][k].
//
// For each element v, for each kk:
//   len = max(best[v][kk] + 1, (kk>0 ? global[kk-1] + 1 : 1))

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumLength(nums []int, k int) int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return 0
	}
	// bestSame[val][kk] = max length ending with val using at most kk diffs.
  // HashMap: O(1) lookup
	bestSame := make(map[int][]int)
	// global[kk] = overall max length using at most kk diffs.
  // Alokasi slice
	global := make([]int, k+1)

	for _, v := range nums {
		if bestSame[v] == nil {
			bestSame[v] = make([]int, k+1)
		}
		row := bestSame[v]
		// Use temporary slice to avoid using updated values within the same
		// iteration (we need the state before processing this element).
  // Alokasi slice
		newBest := make([]int, k+1)
		copy(newBest, row)

		for kk := 0; kk <= k; kk++ {
			cur := 1
			if row[kk] > 0 {
				cur = max(cur, row[kk]+1)
			}
			if kk > 0 && global[kk-1] > 0 {
				cur = max(cur, global[kk-1]+1)
			}
			newBest[kk] = max(newBest[kk], cur)
			global[kk] = max(global[kk], cur)
		}
		bestSame[v] = newBest
	}

	ans := 0
	for kk := 0; kk <= k; kk++ {
		ans = max(ans, global[kk])
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 1, 1, 3}, 2)) // expect 4 (e.g. [1,2,1,1])
	fmt.Println(maximumLength([]int{1, 2, 3, 4, 5}, 1)) // expect 2
}
```
