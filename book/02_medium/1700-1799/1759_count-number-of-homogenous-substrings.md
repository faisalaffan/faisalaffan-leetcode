# 1759 — Count Number Of Homogenous Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countHomogenous(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1759: Count Number of Homogenous Substrings
// https://leetcode.com/problems/count-number-of-homogenous-substrings/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

const mod = 1_000_000_007

func countHomogenous(s string) int {
	result := 0
	count := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		result = (result + count) % mod
	}
	return result
}

func main() {
	fmt.Println(countHomogenous("abbcccaa")) // Expected: 13
	fmt.Println(countHomogenous("xy"))        // Expected: 2
	fmt.Println(countHomogenous("zzzzz"))     // Expected: 15
}
```
