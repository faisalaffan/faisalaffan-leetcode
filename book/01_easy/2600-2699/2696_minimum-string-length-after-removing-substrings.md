# 2696 — Minimum String Length After Removing Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinimumStringLengthAfterRemovingSubstrings(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2696: Minimum String Length After Removing Substrings
// https://leetcode.com/problems/minimum-string-length-after-removing-substrings/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinimumStringLengthAfterRemovingSubstrings("ABFCACDB"))
	fmt.Println(MinimumStringLengthAfterRemovingSubstrings("ACBBD"))
}

func MinimumStringLengthAfterRemovingSubstrings(s string) int {
	stack := []rune{}
	for _, c := range s {
		if len(stack) > 0 &&
			((stack[len(stack)-1] == 'A' && c == 'B') ||
				(stack[len(stack)-1] == 'C' && c == 'D')) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack)
}
```
