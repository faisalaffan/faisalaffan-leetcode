# 1528 — Shuffle String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func restoreString(s string, indices []int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1528: Shuffle String
// https://leetcode.com/problems/shuffle-string/
// Difficulty: Easy
//
// LeetCode submission: func restoreString(s string, indices []int) string

import "fmt"

func main() {
	fmt.Println(ShuffleString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})) // "leetcode"
	fmt.Println(ShuffleString("abc", []int{0, 1, 2}))                     // "abc"
}

// Time: O(n), Space: O(n)
func ShuffleString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i, idx := range indices {
		res[idx] = s[i]
	}
	return string(res)
}
```
