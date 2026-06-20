# 2953 — Count Complete Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countCompleteSubstrings(word string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2953: Count Complete Substrings
// https://leetcode.com/problems/count-complete-substrings/
// Difficulty: Hard
//
// A substring is complete if:
//   1. Every character appears exactly k times.
//   2. For any two adjacent chars, |ord(c1)-ord(c2)| <= 2.
//
// Split string at positions where adjacent diff > 2, then for each segment
// try window sizes = i*k for i = 1..26. Use frequency-of-frequencies to
// check completeness in O(1) per window slide.

import (
	"fmt"
)

func countCompleteSubstrings(word string, k int) int {
	n := len(word)

	countInSegment := func(s string) int {
		m := len(s)
		res := 0
		for distinct := 1; distinct <= 26; distinct++ {
			winLen := distinct * k
			if winLen > m {
				break
			}

  // Alokasi slice
			cnt := make([]int, 26)
  // Alokasi slice
			freq := make([]int, m+1)
			freq[0] = 26

			for i := 0; i < winLen; i++ {
				idx := s[i] - 'a'
				freq[cnt[idx]]--
				cnt[idx]++
				freq[cnt[idx]]++
			}
			if freq[k] == distinct {
				res++
			}

			for i := winLen; i < m; i++ {
				right := s[i] - 'a'
				freq[cnt[right]]--
				cnt[right]++
				freq[cnt[right]]++

				left := s[i-winLen] - 'a'
				freq[cnt[left]]--
				cnt[left]--
				freq[cnt[left]]++

				if freq[k] == distinct {
					res++
				}
			}
		}
		return res
	}

	ans := 0
	start := 0
	for i := 1; i < n; i++ {
		diff := int(word[i]) - int(word[i-1])
		if diff < 0 {
			diff = -diff
		}
		if diff > 2 {
			ans += countInSegment(word[start:i])
			start = i
		}
	}
	ans += countInSegment(word[start:])
	return ans
}

func main() {
	// Example: "igigee", k=2 -> 3
	fmt.Println(countCompleteSubstrings("igigee", 2))

	// Edge cases
	fmt.Println(countCompleteSubstrings("aa", 2))
	fmt.Println(countCompleteSubstrings("abc", 1))
	fmt.Println(countCompleteSubstrings("ab", 1))
	fmt.Println(countCompleteSubstrings("aaaa", 2))
}
```
