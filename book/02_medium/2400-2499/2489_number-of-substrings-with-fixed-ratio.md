# 2489 — Number Of Substrings With Fixed Ratio

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func fixedRatio(s string, num1 int, num2 int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2489: Number of Substrings With Fixed Ratio
// https://leetcode.com/problems/number-of-substrings-with-fixed-ratio/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Count substrings where count('0') : count('1') = num1 : num2.
// Transform: track (num2 * cnt0 - num1 * cnt1), count equal values.

import "fmt"

func main() {
	fmt.Println(fixedRatio("01001", 2, 3)) // 2
	fmt.Println(fixedRatio("0000", 1, 1))  // 0
}

func fixedRatio(s string, num1 int, num2 int) int64 {
  // HashMap: O(1) lookup
	prefix := make(map[int]int64)
	prefix[0] = 1
	var cnt0, cnt1 int64
	var ans int64

	for _, ch := range s {
		if ch == '0' {
			cnt0++
		} else {
			cnt1++
		}
		key := num2*int(cnt0) - num1*int(cnt1)
		ans += prefix[key]
		prefix[key]++
	}
	return ans
}
```
