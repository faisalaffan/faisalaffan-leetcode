# 1513 — Number Of Substrings With Only 1S

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func NumSub(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1513: Number of Substrings With Only 1s
// https://leetcode.com/problems/number-of-substrings-with-only-1s/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumSub("0110111"))
	fmt.Println(NumSub("101"))
	fmt.Println(NumSub("111111"))
}

func NumSub(s string) int {
	// Time: O(N), Space: O(1)
	const mod = 1_000_000_007

	count := 0
	consecutive := 0

	for _, ch := range s {
		if ch == '1' {
			consecutive++
			// Each new 1 adds 'consecutive' new substrings ending at this position
			count = (count + consecutive) % mod
		} else {
			consecutive = 0
		}
	}

	return count
}
```
