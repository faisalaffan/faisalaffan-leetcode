# 0076 — Minimum Window Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minWindow(s string, t string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #76: Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("76. Minimum Window Substring")
	fmt.Printf("s=ADOBECODEBANC, t=ABC -> %q (expected BANC)\n", minWindow("ADOBECODEBANC", "ABC"))
	fmt.Printf("s=a, t=a -> %q (expected a)\n", minWindow("a", "a"))
	fmt.Printf("s=a, t=aa -> %q (expected empty)\n", minWindow("a", "aa"))
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	need := [128]int{}
  // Linear scan O(n)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := [128]int{}
	haveCount := 0
	needCount := 0
	for _, v := range need {
		if v > 0 {
			needCount++
		}
	}

	left := 0
	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		have[c]++
		if have[c] == need[c] {
			haveCount++
		}

		for haveCount == needCount {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			lc := s[left]
			if have[lc] == need[lc] {
				haveCount--
			}
			have[lc]--
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}
	return s[start : start+minLen]
}
```
