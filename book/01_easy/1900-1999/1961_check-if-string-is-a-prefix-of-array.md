# 1961 — Check If String Is A Prefix Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckIfStringIsAPrefixOfArray(s string, words []string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1961: Check If String Is a Prefix of Array
// https://leetcode.com/problems/check-if-string-is-a-prefix-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfStringIsAPrefixOfArray("iloveleetcode", []string{"i", "love", "leetcode", "apples"})) // true
	fmt.Println(CheckIfStringIsAPrefixOfArray("iloveleetcode", []string{"apples", "i", "love", "leetcode"})) // false
}

// Time: O(n), Space: O(1)
func CheckIfStringIsAPrefixOfArray(s string, words []string) bool {
	i := 0
	for _, w := range words {
		if i >= len(s) {
			break
		}
		for j := 0; j < len(w); j++ {
			if i >= len(s) || s[i] != w[j] {
				return false
			}
			i++
		}
	}
	return i == len(s)
}
```
