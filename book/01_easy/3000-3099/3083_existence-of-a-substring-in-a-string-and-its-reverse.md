# 3083 — Existence Of A Substring In A String And Its Reverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func ExistenceOfASubstringInAStringAndItsReverse(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3083: Existence of a Substring in a String and Its Reverse
// https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isSubstringPresent
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("leetcode")) // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcba"))   // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcd"))    // false
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: isSubstringPresent
func ExistenceOfASubstringInAStringAndItsReverse(s string) bool {
	// Build set of all substrings of length 2
  // Membuat map (HashMap) — pencarian O(1)
	substrings := make(map[string]bool)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s)-1; i++ {
		substrings[s[i:i+2]] = true
	}

	// Check reverse for any of those substrings
	for i := len(s) - 1; i > 0; i-- {
		if substrings[string(s[i])+string(s[i-1])] {
			return true
		}
	}
	return false
}
```
