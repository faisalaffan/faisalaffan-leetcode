# 0917 — Reverse Only Letters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func reverseOnlyLetters(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #917: Reverse Only Letters
// https://leetcode.com/problems/reverse-only-letters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))      // "dc-ba"
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj")) // "j-Ih-gfE-dCba"
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!")) // "Qedo1ct-eeLg=ntse-T!"
}

// reverseOnlyLetters reverses only the letters in the string, keeping non-letters in place.
// Time: O(n). Space: O(n).
func reverseOnlyLetters(s string) string {
	b := []byte(s)
	l, r := 0, len(b)-1
	for l < r {
		if !isLetter(b[l]) {
			l++
			continue
		}
		if !isLetter(b[r]) {
			r--
			continue
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
```
