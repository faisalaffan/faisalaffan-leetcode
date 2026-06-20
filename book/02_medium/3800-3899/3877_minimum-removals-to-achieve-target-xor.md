# 3877 — Minimum Removals To Achieve Target Xor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumRemovalsToAchieveTargetXor(nums []int, target int) int
```

> **💡 Hint:** DP tracking max selectable elements to achieve each XOR value.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(N * 2^M)  
**Kompleksitas Ruang:** O(2^M) where M = max bit length (14)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3877: Minimum Removals to Achieve Target XOR
// https://leetcode.com/problems/minimum-removals-to-achieve-target-xor/
// Difficulty: Medium
// Time: O(N * 2^M) | Space: O(2^M) where M = max bit length (14)
// Approach: DP tracking max selectable elements to achieve each XOR value.
// Answer = len(nums) - maxElementsForTarget (or -1 if unreachable).

import "fmt"

func MinimumRemovalsToAchieveTargetXor(nums []int, target int) int {
	maxXor := 1
	for _, v := range nums {
		for maxXor <= v {
			maxXor <<= 1
		}
	}
	if maxXor <= target {
		for maxXor <= target {
			maxXor <<= 1
		}
	}

	// dp[x] = max elements selectable to achieve XOR x
  // Alokasi slice integer
	dp := make([]int, maxXor)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for _, v := range nums {
  // Alokasi slice integer
		ndp := make([]int, maxXor)
		copy(ndp, dp)
		for x := 0; x < maxXor; x++ {
			if dp[x] >= 0 {
				nx := x ^ v
				if dp[x]+1 > ndp[nx] {
					ndp[nx] = dp[x] + 1
				}
			}
		}
		dp = ndp
	}

	if dp[target] < 0 {
		return -1
	}
	return len(nums) - dp[target]
}

func main() {
	// Example 1
	fmt.Println(MinimumRemovalsToAchieveTargetXor([]int{1, 2, 3}, 2)) // Expected: 1

	// Example 2
	fmt.Println(MinimumRemovalsToAchieveTargetXor([]int{2, 4}, 1)) // Expected: -1

	// Example 3
	fmt.Println(MinimumRemovalsToAchieveTargetXor([]int{7}, 7)) // Expected: 0
}
```
