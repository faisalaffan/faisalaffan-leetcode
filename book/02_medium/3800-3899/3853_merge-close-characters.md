# 3853 — Merge Close Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MergeCloseCharacters(s string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N^2)  |  **Ruang:** O(N)


## 💻 Solusi Go

```go
package main

// LeetCode #3853: Merge Close Characters
// https://leetcode.com/problems/merge-close-characters/
// Difficulty: Medium
// Time: O(N^2) | Space: O(N)
// Approach: Simulate merging. Track last position of each character.
// When a char appears within distance k of its last occurrence, skip it.

import "fmt"

func MergeCloseCharacters(s string, k int) string {
	result := make([]byte, 0)

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		// Check if same char exists in result within distance k
		merge := false
		for j := len(result) - 1; j >= 0 && len(result)-1-j < k; j-- {
			if result[j] == ch {
				merge = true
				break
			}
		}
		if !merge {
			result = append(result, ch)
		}
	}

	return string(result)
}

func main() {
	// Example 1
	fmt.Println(MergeCloseCharacters("abca", 3)) // Expected: "abc"

	// Example 2
	fmt.Println(MergeCloseCharacters("aabca", 2)) // Expected: "abca"

	// Example 3
	fmt.Println(MergeCloseCharacters("yybyzybz", 2)) // Expected: "ybzybz"
}
```
