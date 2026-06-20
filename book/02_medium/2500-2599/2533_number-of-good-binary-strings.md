# 2533 — Number Of Good Binary Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func goodBinaryStrings(minLength int, maxLength int, oneGroup int, zeroGroup int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(maxLen)  
**Kompleksitas Ruang:** O(maxLen)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2533: Number of Good Binary Strings
// https://leetcode.com/problems/number-of-good-binary-strings/
// Difficulty: Medium
// Time: O(maxLen) | Space: O(maxLen)
// DP: dp[i] = number of good strings of length i.
// Recurrence: dp[i] = dp[i-oneGroup] + dp[i-zeroGroup]
// Each maximal block of 1s must have length multiple of oneGroup.
// Each maximal block of 0s must have length multiple of zeroGroup.

import "fmt"

func main() {
	fmt.Println(goodBinaryStrings(2, 3, 1, 2)) // 5
	fmt.Println(goodBinaryStrings(3, 3, 1, 1)) // 8
}

const MOD = 1000000007

func goodBinaryStrings(minLength int, maxLength int, oneGroup int, zeroGroup int) int {
  // Alokasi slice integer
	dp := make([]int, maxLength+1)
	dp[0] = 1

	for i := 1; i <= maxLength; i++ {
		if i >= oneGroup {
			dp[i] = (dp[i] + dp[i-oneGroup]) % MOD
		}
		if i >= zeroGroup {
			dp[i] = (dp[i] + dp[i-zeroGroup]) % MOD
		}
	}

	ans := 0
	for i := minLength; i <= maxLength; i++ {
		ans = (ans + dp[i]) % MOD
	}
	return ans
}
```
