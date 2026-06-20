# 2063 — Vowels Of All Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countVowels(word string) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2063: Vowels of All Substrings
// https://leetcode.com/problems/vowels-of-all-substrings/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countVowels(word string) int64 {
	n := len(word)
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	var result int64 = 0

	for i := 0; i < n; i++ {
		if vowels[word[i]] {
			// Number of substrings containing word[i]
			// = (i+1) * (n-i)
			result += int64(i+1) * int64(n-i)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countVowels("aba"))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", countVowels("abc"))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", countVowels("no"))
	// Expected: 0
}
```
