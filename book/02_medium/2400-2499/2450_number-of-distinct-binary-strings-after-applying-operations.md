# 2450 — Number Of Distinct Binary Strings After Applying Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func distinctBinaryStrings(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2450: Number of Distinct Binary Strings After Applying Operations
// https://leetcode.com/problems/number-of-distinct-binary-strings-after-applying-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Count distinct binary strings reachable by inverting any k-length substring.
// This is equivalent to 2^(number_of_free_variables).

import "fmt"

func main() {
	fmt.Println(distinctBinaryStrings("110", 2)) // 4
	fmt.Println(distinctBinaryStrings("10110", 5)) // 2
}

const MOD = 1000000007

func distinctBinaryStrings(s string, k int) int {
	n := len(s)
	// Number of reachable strings = 2^(n-k+1) if k > 0
	// Because each of the first n-k+1 bits can be independently flipped
	if k > n {
		return 1
	}
	pow := 1
	for i := 0; i < n-k+1; i++ {
		pow = (pow * 2) % MOD
	}
	return pow
}
```
