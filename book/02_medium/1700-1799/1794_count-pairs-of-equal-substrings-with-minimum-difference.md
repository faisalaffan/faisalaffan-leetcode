# 1794 — Count Pairs Of Equal Substrings With Minimum Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countQuadruples(firstString string, secondString string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1794: Count Pairs of Equal Substrings With Minimum Difference
// https://leetcode.com/problems/count-pairs-of-equal-substrings-with-minimum-difference/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func countQuadruples(firstString string, secondString string) int {
  // Alokasi slice integer
	firstPos := make([]int, 26)
  // Alokasi slice integer
	lastPos := make([]int, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range firstPos {
		firstPos[i] = -1
		lastPos[i] = -1
	}

	for i, ch := range firstString {
		idx := ch - 'a'
		if firstPos[idx] == -1 {
			firstPos[idx] = i
		}
	}
	for i, ch := range secondString {
		idx := ch - 'a'
		lastPos[idx] = i
	}

	minDiff := 1 << 30
	count := 0

	for i := 0; i < 26; i++ {
		if firstPos[i] != -1 && lastPos[i] != -1 {
			diff := firstPos[i] - lastPos[i]
			if diff < minDiff {
				minDiff = diff
				count = 1
			} else if diff == minDiff {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countQuadruples("abcd", "bcd")) // test 1
	fmt.Println(countQuadruples("abc", "abc")) // test 2
	fmt.Println(countQuadruples("abb", "b")) // test 3
}
```
