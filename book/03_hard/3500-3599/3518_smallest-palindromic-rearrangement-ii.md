# 3518 — Smallest Palindromic Rearrangement Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func smallestPalindrome(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3518: Smallest Palindromic Rearrangement II
// https://leetcode.com/problems/smallest-palindromic-rearrangement-ii/
// Difficulty: Hard
//
// Given a string, rearrange it to form the k-th smallest palindrome
// (lexicographically) among all palindromic rearrangements. Return the
// k-th smallest palindrome, or empty string if fewer than k exist.
//
// Approach: Count character frequencies. Only at most one odd count is
// allowed for a palindrome. Generate the k-th by iterating through
// character positions with combinatorics.

import "fmt"

func main() {
	// Example 1
	fmt.Println(smallestPalindrome("aabb", 2))
	// Example 2
	fmt.Println(smallestPalindrome("a", 1))
	// Example 3: no palindrome possible
	fmt.Println(smallestPalindrome("abc", 1))
	// Edge: k > count
	fmt.Println(smallestPalindrome("abba", 10))
	// Edge: longer string
	fmt.Println(smallestPalindrome("aaabbb", 3))
}

func smallestPalindrome(s string, k int) string {
	// Count frequencies
  // Alokasi slice
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Check if palindrome is possible
	oddCount := 0
	oddChar := -1
	for i, c := range freq {
		if c%2 == 1 {
			oddCount++
			oddChar = i
		}
	}
	if oddCount > 1 {
		return ""
	}

	// Build the half (first half of palindrome)
  // Alokasi slice
	half := make([]int, 0)
	for i, c := range freq {
		for j := 0; j < c/2; j++ {
			half = append(half, i)
		}
	}

	m := len(half)
	if m == 0 {
		if k == 1 {
			return string(byte('a' + oddChar))
		}
		return ""
	}

	// Generate k-th permutation of half using factorial number system
	// Count total permutations
  // Alokasi slice
	fact := make([]int, m+1)
	fact[0] = 1
	for i := 1; i <= m; i++ {
		fact[i] = fact[i-1] * i
	}

	total := fact[m]
	// Adjust for duplicate characters
	for _, c := range freq {
		for j := 2; j <= c/2; j++ {
			total /= j
		}
	}

	if k > total {
		return ""
	}

	// Use combinatorial generation (k-th permutation with duplicates)
	var buildHalf func(remaining []int, k int) []int
	buildHalf = func(remaining []int, k int) []int {
		if len(remaining) == 0 {
			return nil
		}
		// Count distinct values and their frequencies
		type charFreq struct {
			char int
			freq int
		}
		var cf []charFreq
  // HashMap: O(1) lookup
		seen := make(map[int]int)
		for _, v := range remaining {
			seen[v]++
		}
		for c, f := range seen {
			cf = append(cf, charFreq{c, f})
		}

		// For each distinct char at this position
		for _, cfi := range cf {
			// Compute permutations of the rest
			permCount := fact[len(remaining)-1]
			for _, cfj := range cf {
				cnt := cfj.freq
				if cfj.char == cfi.char {
					cnt--
				}
				for j := 2; j <= cnt; j++ {
					permCount /= j
				}
			}
			if k <= permCount {
				// Place this char
  // Alokasi slice
				newRemaining := make([]int, 0)
				placed := false
				for _, v := range remaining {
					if !placed && v == cfi.char {
						placed = true
					} else {
						newRemaining = append(newRemaining, v)
					}
				}
				return append([]int{cfi.char}, buildHalf(newRemaining, k)...)
			}
			k -= permCount
		}
		return nil
	}

	permHalf := buildHalf(half, k)
	if permHalf == nil {
		return ""
	}

	// Build full palindrome
	var result []byte
	for _, v := range permHalf {
		result = append(result, byte('a'+v))
	}
	if oddChar >= 0 {
		result = append(result, byte('a'+oddChar))
	}
	for i := len(permHalf) - 1; i >= 0; i-- {
		result = append(result, byte('a'+permHalf[i]))
	}

	return string(result)
}
```
