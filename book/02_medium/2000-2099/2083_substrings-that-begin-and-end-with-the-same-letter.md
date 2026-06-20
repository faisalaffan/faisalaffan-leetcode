# 2083 — Substrings That Begin And End With The Same Letter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSubstrings(s string) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2083: Substrings That Begin and End With the Same Letter
// https://leetcode.com/problems/substrings-that-begin-and-end-with-the-same-letter/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfSubstrings(s string) int64 {
  // Alokasi slice integer
	freq := make([]int64, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	var result int64 = 0
	for _, f := range freq {
		result += f * (f + 1) / 2
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfSubstrings("abc"))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", numberOfSubstrings("abacaba"))
	// Expected: 12

	// Test case 3
	fmt.Println("Test 3:", numberOfSubstrings("aa"))
	// Expected: 3
}
```
