# 3329 — Count Substrings With K Frequency Characters Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numberOfSubstrings(s string, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3329: Count Substrings With K-Frequency Characters II
// https://leetcode.com/problems/count-substrings-with-k-frequency-characters-ii/
// Difficulty: Hard [Paid]
//
// Count substrings where at least one character appears at least k times.
//
// Approach: Sliding window with two pointers. For each right endpoint,
// maintain window [left, right] where no character reaches frequency k.
// All substrings starting at [0, left-1] and ending at right are valid.

import "fmt"

func main() {
	// Example 1
	fmt.Println(numberOfSubstrings("abacb", 2))
	// Example 2
	fmt.Println(numberOfSubstrings("abcde", 1))
	// Example 3
	fmt.Println(numberOfSubstrings("aaaaa", 2))
	// Edge: no valid substrings
	fmt.Println(numberOfSubstrings("abc", 5))
	// Single char repeated
	fmt.Println(numberOfSubstrings("aa", 2))
}

func numberOfSubstrings(s string, k int) int64 {
	freq := [26]int{}
	var total int64
	left := 0

	for right := 0; right < len(s); right++ {
		cur := s[right] - 'a'
		freq[cur]++

		for freq[cur] >= k {
			freq[s[left]-'a']--
			left++
		}

		total += int64(left)
	}

	return total
}
```
