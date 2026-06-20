# 3503 — Longest Palindrome After Substring Concatenation I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestPalindromeAfterSubstringConcatenationI(s string, t string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3503: Longest Palindrome After Substring Concatenation I
// https://leetcode.com/problems/longest-palindrome-after-substring-concatenation-i/
// Difficulty: Medium
// Complexity: O(n^2 * m^2) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", LongestPalindromeAfterSubstringConcatenationI("ab", "ba"))
	// Test case 2
	fmt.Println("Test 2:", LongestPalindromeAfterSubstringConcatenationI("a", "a"))
	// Test case 3
	fmt.Println("Test 3:", LongestPalindromeAfterSubstringConcatenationI("abc", "cba"))
}

func LongestPalindromeAfterSubstringConcatenationI(s string, t string) int {
	isPal := func(str string) bool {
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}

	maxLen := 0
	// Try all substrings of s concatenated with all substrings of t
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			for p := 0; p <= len(t); p++ {
				for q := p; q <= len(t); q++ {
					candidate := s[i:j] + t[p:q]
					if isPal(candidate) && len(candidate) > maxLen {
						maxLen = len(candidate)
					}
				}
			}
		}
	}
	return maxLen
}
```
