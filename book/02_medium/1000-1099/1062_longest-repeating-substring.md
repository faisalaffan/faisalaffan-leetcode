# 1062 — Longest Repeating Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestRepeatingSubstring(s string) int
```

> **💡 Hint:** DP - find longest common prefix between all pairs of suffixes

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2) can be O(n) with optimized DP

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1062: Longest Repeating Substring
// https://leetcode.com/problems/longest-repeating-substring/
// Difficulty: Medium
//
// Approach: DP - find longest common prefix between all pairs of suffixes
// Time: O(n^2)
// Space: O(n^2) can be O(n) with optimized DP

import "fmt"

func main() {
	fmt.Println(longestRepeatingSubstring("abcd"))    // 0
	fmt.Println(longestRepeatingSubstring("abbaba"))  // 2
	fmt.Println(longestRepeatingSubstring("aabcaabdaab")) // 3
}

func longestRepeatingSubstring(s string) int {
	n := len(s)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > result {
					result = dp[i][j]
				}
			}
		}
	}

	return result
}
```
