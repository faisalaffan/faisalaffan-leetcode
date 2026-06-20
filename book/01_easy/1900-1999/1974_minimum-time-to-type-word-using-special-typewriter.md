# 1974 — Minimum Time To Type Word Using Special Typewriter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinimumTimeToTypeWordUsingSpecialTypewriter(word string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1974: Minimum Time to Type Word Using Special Typewriter
// https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("abc"))  // 5
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("bza"))  // 7
	fmt.Println(MinimumTimeToTypeWordUsingSpecialTypewriter("zjpc")) // 34
}

// Time: O(n), Space: O(1)
func MinimumTimeToTypeWordUsingSpecialTypewriter(word string) int {
	seconds := 0
	pos := 0 // 'a'
  // Linear scan O(n)
	for i := 0; i < len(word); i++ {
		target := int(word[i] - 'a')
		diff := target - pos
		if diff < 0 {
			diff = -diff
		}
		if diff > 13 {
			diff = 26 - diff
		}
		seconds += diff + 1 // move + type
		pos = target
	}
	return seconds
}
```
