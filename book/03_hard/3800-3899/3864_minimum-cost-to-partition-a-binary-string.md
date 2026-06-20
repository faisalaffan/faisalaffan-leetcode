# 3864 — Minimum Cost To Partition A Binary String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minCost(s string, encCost int, flatCost int) int64
```

> **💡 Hint:** DP[i] = min cost to partition prefix of length i. For

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3864: Minimum Cost to Partition a Binary String
// https://leetcode.com/problems/minimum-cost-to-partition-a-binary-string/
// Difficulty: Hard
//
// Partition binary string into contiguous substrings. Each substring
// of length L costs encCost if it contains only '0's or only '1's,
// otherwise flatCost. Minimize total cost.
//
// Approach: DP[i] = min cost to partition prefix of length i. For
// each j < i, compute cost of substring s[j:i] and update.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(minCost("0101", 2, 3))
	// Example 2
	fmt.Println(minCost("000", 2, 5))
	// Edge: single char
	fmt.Println(minCost("1", 1, 5))
	// Edge: alternating
	fmt.Println(minCost("101010", 1, 10))
}

func minCost(s string, encCost int, flatCost int) int64 {
	n := len(s)
  // Alokasi slice integer
	dp := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		ones := 0
		zeros := 0
		for j := i - 1; j >= 0; j-- {
			if s[j] == '1' {
				ones++
			} else {
				zeros++
			}
			length := i - j
			var cost int64
			if ones == 0 || zeros == 0 {
				cost = int64(encCost) * int64(length)
			} else {
				cost = int64(flatCost)
			}
			if dp[j]+cost < dp[i] {
				dp[i] = dp[j] + cost
			}
		}
	}
	return dp[n]
}
```
