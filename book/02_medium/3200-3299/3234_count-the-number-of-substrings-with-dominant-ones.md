# 3234 — Count The Number Of Substrings With Dominant Ones

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

**Kompleksitas Waktu:** O(n * sqrt(n))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3234: Count the Number of Substrings With Dominant Ones
// https://leetcode.com/problems/count-the-number-of-substrings-with-dominant-ones/
// Difficulty: Medium
// Time: O(n * sqrt(n)) | Space: O(1)

import (
	"fmt"
	"math"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	ans := 0
	maxZeros := int(math.Sqrt(float64(n)))

	for l := 0; l < n; l++ {
		zeros := 0
		ones := 0
		for r := l; r < n; r++ {
			if s[r] == '0' {
				zeros++
				if zeros > maxZeros {
					break
				}
			} else {
				ones++
			}
			if ones >= zeros*zeros {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubstrings("00011")) // Expected: 5
	fmt.Println(numberOfSubstrings("101"))    // Expected: 3
}
```
