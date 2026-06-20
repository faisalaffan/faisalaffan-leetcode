# 1933 — Check If String Is Decomposable Into Value Equal Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfStringIsDecomposableIntoValueEqualSubstrings(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1933: Check if String Is Decomposable Into Value-Equal Substrings
// https://leetcode.com/problems/check-if-string-is-decomposable-into-value-equal-substrings/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("000111000"))   // false
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("00011111222"))  // true
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("011100022233")) // false
}

// Time: O(n), Space: O(1)
func CheckIfStringIsDecomposableIntoValueEqualSubstrings(s string) bool {
	hasGroupOfTwo := false
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		count := j - i
		if count%3 == 1 {
			return false
		}
		if count%3 == 2 {
			if hasGroupOfTwo {
				return false
			}
			hasGroupOfTwo = true
		}
		i = j
	}
	return hasGroupOfTwo
}
```
