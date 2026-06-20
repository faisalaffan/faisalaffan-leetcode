# 3144 — Minimum Substring Partition Of Equal Character Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSubstringsInPartition(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3144: Minimum Substring Partition of Equal Character Frequency
// https://leetcode.com/problems/minimum-substring-partition-of-equal-character-frequency/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumSubstringsInPartition(s string) int {
	n := len(s)
  // Alokasi slice integer
	dp := make([]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
  // Alokasi slice integer
		freq := make([]int, 26)
		var distinct, maxFreq int
		for j := i - 1; j >= 0; j-- {
			idx := s[j] - 'a'
			if freq[idx] == 0 {
				distinct++
			}
			freq[idx]++
			if freq[idx] > maxFreq {
				maxFreq = freq[idx]
			}

			if maxFreq*distinct == i-j {
				if dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(minimumSubstringsInPartition("fabccddg")) // Expected: 3
	fmt.Println(minimumSubstringsInPartition("abababaccddb")) // Expected: ?
}
```
