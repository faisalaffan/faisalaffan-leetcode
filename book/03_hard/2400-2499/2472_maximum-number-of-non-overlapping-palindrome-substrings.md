# 2472 — Maximum Number Of Non Overlapping Palindrome Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxPalindromes(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2472: Maximum Number of Non-overlapping Palindrome Substrings
// https://leetcode.com/problems/maximum-number-of-non-overlapping-palindrome-substrings/
// Difficulty: Hard
//
// Greedy shortest palindrome. At each position i, find the shortest palindrome
// of length >= k starting at i (max length k+1 suffices since if a longer palindrome
// exists, a shorter one of length k or k+1 also does). Count it and skip past it.

import "fmt"

func main() {
	// Example 1: "abaccdbbd", 3 => 2
	fmt.Println(maxPalindromes("abaccdbbd", 3))
	// Example 2: "adbcda", 2 => 1
	fmt.Println(maxPalindromes("adbcda", 2))
	// Edge: single character, k=1
	fmt.Println(maxPalindromes("a", 1))
	// Edge: no palindrome of length >= k
	fmt.Println(maxPalindromes("ab", 3))
	// Edge: overlaps — shortest palindrome at each position wins
	fmt.Println(maxPalindromes("aaa", 2))
}

func maxPalindromes(s string, k int) int {
	n := len(s)
	count := 0
	i := 0

	for i < n {
		found := false
		// We only need to check up to k+1 because if any palindrome >= k exists,
		// either k or k+1 is a palindrome (by parity).
		limit := k + 1
		if n-i < limit {
			limit = n - i
		}
		for length := k; length <= limit; length++ {
			if isPalindrome(s, i, i+length-1) {
				count++
				i += length
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}

	return count
}

func isPalindrome(s string, l, r int) bool {
  // Two-pointer: gerakkan kiri atau kanan
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
```
