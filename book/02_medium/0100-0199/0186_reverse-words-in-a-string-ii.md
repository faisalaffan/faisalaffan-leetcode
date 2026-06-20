# 0186 — Reverse Words In A String Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func reverseWords(s []byte) `

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1) in-place  |  **Ruang:** O(1) in-place


## 💻 Solusi Go

```go
package main

// LeetCode #186: Reverse Words in a String II
// https://leetcode.com/problems/reverse-words-in-a-string-ii/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1) in-place

import "fmt"

func reverseWords(s []byte) {
	reverse := func(arr []byte, l, r int) {
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	reverse(s, 0, len(s)-1)

	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			reverse(s, start, i-1)
			start = i + 1
		}
	}
}

func main() {
	s1 := []byte("the sky is blue")
	reverseWords(s1)
	fmt.Println(string(s1))

	s2 := []byte("hello world")
	reverseWords(s2)
	fmt.Println(string(s2))

	s3 := []byte("a")
	reverseWords(s3)
	fmt.Println(string(s3))
}
```
