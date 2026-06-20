# 3777 — Minimum Deletions To Make Alternating Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minDeletions(s string, queries [][]int) []int64
```

> **💡 Hint:** For each query, compute longest alternating subsequence

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3777: Minimum Deletions to Make Alternating Substring
// https://leetcode.com/problems/minimum-deletions-to-make-alternating-substring/
// Difficulty: Hard
//
// For each query [l, r], find minimum deletions to make substring
// s[l..r] alternating (no two adjacent characters are the same).
//
// Approach: For each query, compute longest alternating subsequence
// in the substring. Answer = length - LAS length.
// DP tracks longest alternating subsequence ending with each char.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minDeletions("abcabc", [][]int{{0, 5}}))
	// Example 2
	fmt.Println(minDeletions("aabbaa", [][]int{{0, 5}, {1, 4}}))
	// Edge: single char
	fmt.Println(minDeletions("a", [][]int{{0, 0}}))
	// Edge: already alternating
	fmt.Println(minDeletions("abab", [][]int{{0, 3}}))
}

func minDeletions(s string, queries [][]int) []int64 {
	n := len(s)
  // Alokasi slice integer
	ans := make([]int64, len(queries))

	for qi, q := range queries {
		l, r := q[0], q[1]
		if l >= n || r >= n || l > r {
			ans[qi] = 0
			continue
		}
		length := r - l + 1
		if length <= 1 {
			ans[qi] = 0
			continue
		}

		// DP for longest alternating subsequence
  // Alokasi slice integer
		dp := make([]int, 26)
		for i := l; i <= r; i++ {
			c := int(s[i] - 'a')
			bestPrev := 0
			for pc := 0; pc < 26; pc++ {
				if pc != c && dp[pc] > bestPrev {
					bestPrev = dp[pc]
				}
			}
			if bestPrev+1 > dp[c] {
				dp[c] = bestPrev + 1
			}
		}

		// Find max LAS length
		lasLen := 0
		for _, v := range dp {
			if v > lasLen {
				lasLen = v
			}
		}

		ans[qi] = int64(length - lasLen)
	}

	return ans
}
```
