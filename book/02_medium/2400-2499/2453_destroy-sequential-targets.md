# 2453 — Destroy Sequential Targets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func destroyTargets(nums []int, space int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2453: Destroy Sequential Targets
// https://leetcode.com/problems/destroy-sequential-targets/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Group nums by nums[i] % space. The group with max size gives max targets.
// Pick smallest nums[i] from that group.

import "fmt"

func main() {
	fmt.Println(destroyTargets([]int{3, 7, 8, 1, 1, 5}, 2)) // 1
	fmt.Println(destroyTargets([]int{1, 3, 5, 2, 4, 6}, 2)) // 1
}

func destroyTargets(nums []int, space int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
  // HashMap: O(1) lookup
	minVal := make(map[int]int)

	for _, v := range nums {
		rem := v % space
		freq[rem]++
		if _, ok := minVal[rem]; !ok || v < minVal[rem] {
			minVal[rem] = v
		}
	}

	maxFreq, ans := 0, 0
	for rem, f := range freq {
		if f > maxFreq || (f == maxFreq && minVal[rem] < ans) {
			maxFreq = f
			ans = minVal[rem]
		}
	}
	return ans
}
```
