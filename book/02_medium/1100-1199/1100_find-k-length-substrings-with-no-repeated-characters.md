# 1100 — Find K Length Substrings With No Repeated Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numKLenSubstrNoRepeats(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sliding Window

**Waktu:** O(n)  |  **Ruang:** O(26) = O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sliding Window** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1100: Find K-Length Substrings With No Repeated Characters
// https://leetcode.com/problems/find-k-length-substrings-with-no-repeated-characters/
// Difficulty: Medium
//
// Approach: Sliding window with frequency array
// Time: O(n)
// Space: O(26) = O(1)

import "fmt"

func main() {
	fmt.Println(numKLenSubstrNoRepeats("havefunonleetcode", 5)) // 6
	fmt.Println(numKLenSubstrNoRepeats("home", 5))              // 0
}

func numKLenSubstrNoRepeats(s string, k int) int {
	if k > len(s) || k > 26 {
		return 0
	}

  // Alokasi slice
	freq := make([]int, 26)
	duplicates := 0
	result := 0

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		if freq[s[i]-'a'] == 2 {
			duplicates++
		}

		if i >= k {
			freq[s[i-k]-'a']--
			if freq[s[i-k]-'a'] == 1 {
				duplicates--
			}
		}

		if i >= k-1 && duplicates == 0 {
			result++
		}
	}

	return result
}
```
