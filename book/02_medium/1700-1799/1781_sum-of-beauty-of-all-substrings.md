# 1781 — Sum Of Beauty Of All Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func beautySum(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(26)  
**Kompleksitas Ruang:** O(26)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1781: Sum of Beauty of All Substrings
// https://leetcode.com/problems/sum-of-beauty-of-all-substrings/
// Difficulty: Medium
// Time: O(n^2), Space: O(26)

import "fmt"

func beautySum(s string) int {
	n := len(s)
	result := 0

	for i := 0; i < n; i++ {
  // Alokasi slice integer
		count := make([]int, 26)
		for j := i; j < n; j++ {
			count[s[j]-'a']++
			minFreq, maxFreq := n, 0
			for _, f := range count {
				if f > 0 {
					if f < minFreq {
						minFreq = f
					}
					if f > maxFreq {
						maxFreq = f
					}
				}
			}
			result += maxFreq - minFreq
		}
	}
	return result
}

func main() {
	fmt.Println(beautySum("aabcb")) // Expected: 5
	fmt.Println(beautySum("aabcbaa")) // Expected: 17
	fmt.Println(beautySum("x")) // Expected: 0
}
```
