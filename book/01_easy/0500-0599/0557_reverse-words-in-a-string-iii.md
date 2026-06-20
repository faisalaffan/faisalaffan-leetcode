# 0557 — Reverse Words In A String Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ReverseWordsInAStringIii(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #557: Reverse Words in a String III
// https://leetcode.com/problems/reverse-words-in-a-string-iii/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReverseWordsInAStringIii(s string) string {
	b := []byte(s)
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			for lo, hi := start, i-1; lo < hi; lo, hi = lo+1, hi-1 {
				b[lo], b[hi] = b[hi], b[lo]
			}
			start = i + 1
		}
	}
	return string(b)
}

func main() {
	fmt.Println(ReverseWordsInAStringIii("Let's take LeetCode contest"))
	fmt.Println(ReverseWordsInAStringIii("Mr Ding"))
}
```
