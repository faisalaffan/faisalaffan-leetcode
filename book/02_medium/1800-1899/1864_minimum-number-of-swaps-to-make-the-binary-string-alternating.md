# 1864 — Minimum Number Of Swaps To Make The Binary String Alternating

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinSwaps(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1864: Minimum Number of Swaps to Make the Binary String Alternating
// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSwaps("111000"))
	fmt.Println(MinSwaps("010"))
	fmt.Println(MinSwaps("1110"))
}

// Time: O(n), Space: O(1)
func MinSwaps(s string) int {
	n := len(s)
	ones := 0
	zeros := 0
	for _, c := range s {
		if c == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if abs(ones-zeros) > 1 {
		return -1
	}

	// Count mismatches when starting with '0' and starting with '1'
	mismatch0 := 0 // pattern: 010101...
	mismatch1 := 0 // pattern: 101010...
	for i, c := range s {
		if i%2 == 0 {
			if c == '1' {
				mismatch0++
			} else {
				mismatch1++
			}
		} else {
			if c == '0' {
				mismatch0++
			} else {
				mismatch1++
			}
		}
	}

	// Each swap fixes 2 mismatches
	if n%2 == 0 {
		return min(mismatch0/2, mismatch1/2)
	}
	// For odd length, only one pattern is valid
	if zeros > ones {
		return mismatch0 / 2
	}
	return mismatch1 / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
