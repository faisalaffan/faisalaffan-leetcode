# 3557 — Find Maximum Number Of Non Intersecting Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindMaximumNumberOfNonIntersectingSubstrings(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3557: Find Maximum Number of Non Intersecting Substrings
// https://leetcode.com/problems/find-maximum-number-of-non-intersecting-substrings/
// Difficulty: Medium
// Complexity: O(n^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindMaximumNumberOfNonIntersectingSubstrings("abcba"))
	// Test case 2
	fmt.Println("Test 2:", FindMaximumNumberOfNonIntersectingSubstrings("abac"))
	// Test case 3
	fmt.Println("Test 3:", FindMaximumNumberOfNonIntersectingSubstrings("a"))
}

func FindMaximumNumberOfNonIntersectingSubstrings(s string) int {
	n := len(s)
	count := 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		// Try to find a substring starting at i that doesn't intersect with used ones
		for j := i + 1; j <= n; j++ {
			overlap := false
			for k := i; k < j; k++ {
				if used[k] {
					overlap = true
					break
				}
			}
			if !overlap {
				// Check if substring s[i:j] is valid (e.g., palindrome check)
				isPal := true
				for l, r := i, j-1; l < r; l, r = l+1, r-1 {
					if s[l] != s[r] {
						isPal = false
						break
					}
				}
				if isPal {
					for k := i; k < j; k++ {
						used[k] = true
					}
					count++
					break
				}
			}
		}
	}
	return count
}
```
