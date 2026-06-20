# 1370 — Increasing Decreasing String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func sortString(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1370: Increasing Decreasing String
// https://leetcode.com/problems/increasing-decreasing-string/
// Difficulty: Easy
//
// LeetCode submission: func sortString(s string) string

import "fmt"

func main() {
	fmt.Println(IncreasingDecreasingString("aaaabbbbcccc")) // "abccbaabccba"
	fmt.Println(IncreasingDecreasingString("rat"))          // "art"
	fmt.Println(IncreasingDecreasingString("leetcode"))     // "cdelotee"
}

// Time: O(n), Space: O(n)
func IncreasingDecreasingString(s string) string {
  // Alokasi slice
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	res := make([]byte, 0, len(s))
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if freq[i] > 0 {
				res = append(res, byte('a'+i))
				freq[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if freq[i] > 0 {
				res = append(res, byte('a'+i))
				freq[i]--
			}
		}
	}
	return string(res)
}
```
