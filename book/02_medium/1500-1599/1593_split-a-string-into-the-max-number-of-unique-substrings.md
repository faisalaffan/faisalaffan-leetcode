# 1593 — Split A String Into The Max Number Of Unique Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxUniqueSplit(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(2^N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1593: Split a String Into the Max Number of Unique Substrings
// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxUniqueSplit("ababccc"))
	fmt.Println(MaxUniqueSplit("aba"))
	fmt.Println(MaxUniqueSplit("aa"))
}

func MaxUniqueSplit(s string) int {
	// Time: O(2^N), Space: O(N)
  // Membuat map (HashMap) — pencarian O(1)
	used := make(map[string]bool)
	maxCount := 0

	var backtrack func(start int, count int)
	backtrack = func(start int, count int) {
		if start == len(s) {
			if count > maxCount {
				maxCount = count
			}
			return
		}

		// Pruning: remaining chars <= max possible new substrings
		if count+(len(s)-start) <= maxCount {
			return
		}

		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if !used[sub] {
				used[sub] = true
				backtrack(end, count+1)
				used[sub] = false
			}
		}
	}

	backtrack(0, 0)
	return maxCount
}
```
