# 1358 — Number Of Substrings Containing All Three Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSubstrings(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = length of string  
**Kompleksitas Ruang:** O(1) - fixed array of 3

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1358: Number of Substrings Containing All Three Characters
// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numberOfSubstrings("abcabc")) // 10

	// Test case 2
	fmt.Println(numberOfSubstrings("aaacb")) // 3

	// Test case 3
	fmt.Println(numberOfSubstrings("abc")) // 1
}

// Time: O(n) where n = length of string
// Space: O(1) - fixed array of 3
func numberOfSubstrings(s string) int {
	count := 0
	lastPos := [3]int{-1, -1, -1}

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		lastPos[s[i]-'a'] = i
		// Find the minimum of the last positions of 'a', 'b', 'c'
		// This is the leftmost boundary of the substring containing all three
		minPos := lastPos[0]
		if lastPos[1] < minPos {
			minPos = lastPos[1]
		}
		if lastPos[2] < minPos {
			minPos = lastPos[2]
		}
		// If all three have been seen, count substrings ending at i
		if minPos >= 0 {
			count += minPos + 1
		}
	}

	return count
}
```
