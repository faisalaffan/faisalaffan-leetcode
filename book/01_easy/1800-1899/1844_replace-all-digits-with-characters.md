# 1844 — Replace All Digits With Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ReplaceDigits(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1844: Replace All Digits with Characters
// https://leetcode.com/problems/replace-all-digits-with-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func ReplaceDigits(s string) string {
	result := []byte(s)
	for i := 1; i < len(s); i += 2 {
		result[i] = s[i-1] + (s[i] - '0')
	}
	return string(result)
}

func main() {
	fmt.Println(ReplaceDigits("a1c1e1"))
	fmt.Println(ReplaceDigits("a1b2c3d4e"))
}
```
