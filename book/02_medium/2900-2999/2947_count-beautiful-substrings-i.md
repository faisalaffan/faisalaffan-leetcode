# 2947 — Count Beautiful Substrings I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func beautifulSubstrings(s string, k int) (ans int)`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2947: Count Beautiful Substrings I
// https://leetcode.com/problems/count-beautiful-substrings-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(beautifulSubstrings("baeyh", 2))
	fmt.Println(beautifulSubstrings("ab", 1))
	fmt.Println(beautifulSubstrings("a", 1))
}

func beautifulSubstrings(s string, k int) (ans int) {
	n := len(s)
	vowels := [26]bool{}
	for _, c := range "aeiou" {
		vowels[c-'a'] = true
	}
	for i := 0; i < n; i++ {
		v := 0
		for j := i; j < n; j++ {
			if vowels[s[j]-'a'] {
				v++
			}
			c := j - i + 1 - v
			if v == c && v*c%k == 0 {
				ans++
			}
		}
	}
	return
}
```
