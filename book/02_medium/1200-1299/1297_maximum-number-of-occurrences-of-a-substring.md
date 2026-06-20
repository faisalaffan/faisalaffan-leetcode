# 1297 — Maximum Number Of Occurrences Of A Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * minSize)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1297: Maximum Number of Occurrences of a Substring
// https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/
// Difficulty: Medium

// Find max occurrences of any substring meeting constraints:
// - size between 1 and maxLetters distinct characters
// - substring size = minSize (longer substrings are less frequent)

// Time: O(n * minSize)
// Space: O(n)

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[string]int)
	maxOccur := 0

	for i := 0; i+minSize <= len(s); i++ {
		sub := s[i : i+minSize]
		// Count distinct chars
  // Membuat map (HashMap) — pencarian O(1)
		chars := make(map[byte]bool)
		for j := 0; j < len(sub); j++ {
			chars[sub[j]] = true
		}
		if len(chars) <= maxLetters {
			count[sub]++
			if count[sub] > maxOccur {
				maxOccur = count[sub]
			}
		}
	}

	return maxOccur
}

func main() {
	fmt.Printf("%d (expected: 2)\n", maxFreq("aababcaab", 2, 3, 4))
	fmt.Printf("%d (expected: 2)\n", maxFreq("aaaa", 1, 3, 3))
	fmt.Printf("%d (expected: 3)\n", maxFreq("aabcabcab", 2, 3, 3))
}
```
