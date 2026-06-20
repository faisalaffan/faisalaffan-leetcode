# 2451 — Odd String Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func differenceArray(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2451: Odd String Difference
// https://leetcode.com/problems/odd-string-difference/
// Difficulty: Easy
// Time O(n * m) | Space O(n)

import "fmt"

func main() {
	fmt.Println(OddStringDifference([]string{"adc", "wzy", "abc"})) // "abc"
	fmt.Println(OddStringDifference([]string{"aaa", "bob", "ccc", "ddd"})) // "bob"
}

func differenceArray(s string) string {
	diff := make([]byte, len(s)-1)
	for i := 1; i < len(s); i++ {
		diff[i-1] = s[i] - s[i-1]
	}
	return string(diff)
}

func OddStringDifference(words []string) string {
	diff0 := differenceArray(words[0])
	diff1 := differenceArray(words[1])

	if string(diff0) != string(diff1) {
		diff2 := differenceArray(words[2])
		if string(diff0) == string(diff2) {
			return words[1]
		}
		return words[0]
	}

	for i := 2; i < len(words); i++ {
		if string(differenceArray(words[i])) != string(diff0) {
			return words[i]
		}
	}
	return ""
}
```
