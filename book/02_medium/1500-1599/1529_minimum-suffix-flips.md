# 1529 — Minimum Suffix Flips

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinFlips(target string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1529: Minimum Suffix Flips
// https://leetcode.com/problems/minimum-suffix-flips/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinFlips("10111"))
	fmt.Println(MinFlips("101"))
	fmt.Println(MinFlips("00000"))
}

func MinFlips(target string) int {
	// Time: O(N), Space: O(1)
	// Count transitions from 0 to 1 or 1 to 0
	flips := 0
	curr := byte('0') // current state of flipped prefix

  // Linear scan O(n)
	for i := 0; i < len(target); i++ {
		if target[i] != curr {
			flips++
			curr = target[i]
		}
	}

	return flips
}
```
