# 0681 — Next Closest Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func nextClosestTime(time string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(1) since there are at most 4^4 = 256 combinations  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #681: Next Closest Time
// https://leetcode.com/problems/next-closest-time/
// Difficulty: Medium [Paid]
// Time: O(1) since there are at most 4^4 = 256 combinations
// Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(nextClosestTime("19:34"))
	fmt.Println(nextClosestTime("23:59"))
	fmt.Println(nextClosestTime("13:55"))
}

func nextClosestTime(time string) string {
  // HashMap: O(1) lookup
	digits := make(map[byte]bool)
  // Linear scan O(n)
	for i := 0; i < len(time); i++ {
		if time[i] != ':' {
			digits[time[i]] = true
		}
	}

	hours, _ := strconv.Atoi(time[:2])
	minutes, _ := strconv.Atoi(time[3:])

	current := hours*60 + minutes

	for elapsed := 1; elapsed <= 24*60; elapsed++ {
		t := (current + elapsed) % (24 * 60)
		h := t / 60
		m := t % 60

		hs := fmt.Sprintf("%02d%02d", h, m)
		valid := true
		for i := 0; i < 4; i++ {
			if !digits[hs[i]] {
				valid = false
				break
			}
		}

		if valid {
			return fmt.Sprintf("%02d:%02d", h, m)
		}
	}

	return ""
}
```
