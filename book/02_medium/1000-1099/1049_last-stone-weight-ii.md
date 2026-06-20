# 1049 — Last Stone Weight Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func lastStoneWeightII(stones []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n * sum)  |  **Ruang:** O(sum)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1049: Last Stone Weight II
// https://leetcode.com/problems/last-stone-weight-ii/
// Difficulty: Medium
//
// Approach: DP - subset sum. Partition stones into two groups.
//           Minimize |sum - 2*subsetSum|.
// Time: O(n * sum)
// Space: O(sum)

import "fmt"

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1})) // 1
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40})) // 5
}

func lastStoneWeightII(stones []int) int {
	total := 0
	for _, s := range stones {
		total += s
	}

	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, s := range stones {
		for j := target; j >= s; j-- {
			if dp[j-s] {
				dp[j] = true
			}
		}
	}

	for j := target; j >= 0; j-- {
		if dp[j] {
			return total - 2*j
		}
	}

	return total
}
```
