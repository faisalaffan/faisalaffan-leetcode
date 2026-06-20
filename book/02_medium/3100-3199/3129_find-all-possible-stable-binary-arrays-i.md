# 3129 — Find All Possible Stable Binary Arrays I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfStableArrays(zero int, one int, limit int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(zero * one * limit)  
**Kompleksitas Ruang:** O(zero * one * 2)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3129: Find All Possible Stable Binary Arrays I
// https://leetcode.com/problems/find-all-possible-stable-binary-arrays-i/
// Difficulty: Medium
// Time: O(zero * one * limit) | Space: O(zero * one * 2)

import "fmt"

func numberOfStableArrays(zero int, one int, limit int) int {
	const mod = 1_000_000_007
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][2]int, zero+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][2]int, one+1)
	}

	for i := 0; i <= zero; i++ {
		for j := 0; j <= one; j++ {
			if i == 0 && j == 0 {
				dp[i][j][0] = 1
				dp[i][j][1] = 1
				continue
			}
			if i > 0 {
				dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
				if i > limit {
					dp[i][j][0] -= dp[i-limit-1][j][1]
				}
				dp[i][j][0] %= mod
				if dp[i][j][0] < 0 {
					dp[i][j][0] += mod
				}
			}
			if j > 0 {
				dp[i][j][1] = dp[i][j-1][0] + dp[i][j-1][1]
				if j > limit {
					dp[i][j][1] -= dp[i][j-limit-1][0]
				}
				dp[i][j][1] %= mod
				if dp[i][j][1] < 0 {
					dp[i][j][1] += mod
				}
			}
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}

func main() {
	fmt.Println(numberOfStableArrays(1, 1, 2)) // Expected: 2
	fmt.Println(numberOfStableArrays(1, 2, 1)) // Expected: 1
	fmt.Println(numberOfStableArrays(3, 1, 1)) // Expected: 2
}
```
