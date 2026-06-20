# 0848 — Shifting Letters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ShiftingLetters(s string, shifts []int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #848: Shifting Letters
// https://leetcode.com/problems/shifting-letters/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ShiftingLetters("abc", []int{3, 5, 9}))
	fmt.Println(ShiftingLetters("aaa", []int{1, 2, 3}))
	fmt.Println(ShiftingLetters("z", []int{52}))
}

// Time: O(n) | Space: O(n)
func ShiftingLetters(s string, shifts []int) string {
	n := len(s)
	// Calculate suffix sum of shifts
	for i := n - 2; i >= 0; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
	}

	res := []byte(s)
	for i := 0; i < n; i++ {
		res[i] = byte((int(res[i]-'a')+shifts[i])%26 + 'a')
	}

	return string(res)
}
```
