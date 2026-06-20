# 3456 — Find Special Substring Of Length K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindSpecialSubstringOfLengthK(s string, k int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3456: Find Special Substring of Length K
// https://leetcode.com/problems/find-special-substring-of-length-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindSpecialSubstringOfLengthK("aaabaaa", 3))
	fmt.Println(FindSpecialSubstringOfLengthK("abc", 2))
}

// FindSpecialSubstringOfLengthK returns true if there is a substring of length k consisting of a single character, surrounded by different characters (or boundaries).
// Time: O(n). Space: O(1).
func FindSpecialSubstringOfLengthK(s string, k int) bool {
	n := len(s)
	for i := 0; i <= n-k; i++ {
		same := true
		for j := i; j < i+k-1; j++ {
			if s[j] != s[j+1] {
				same = false
				break
			}
		}
		if !same {
			continue
		}
		if i > 0 && s[i-1] == s[i] {
			continue
		}
		if i+k < n && s[i+k] == s[i] {
			continue
		}
		return true
	}
	return false
}
```
