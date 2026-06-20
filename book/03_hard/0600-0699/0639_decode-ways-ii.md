# 0639 — Decode Ways Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numDecodings(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #639: Decode Ways II
// https://leetcode.com/problems/decode-ways-ii/
// Difficulty: Hard

import "fmt"

const mod = 1_000_000_007

func main() {
	// Test cases
	testCases := []struct {
		s    string
		want int
	}{
		{"*", 9},
		{"1*", 18},
		{"**", 96},
		{"*1*1*0", 404},
		{"0", 0},
		{"10", 1},
		{"1", 1},
		{"2", 1},
		{"*0", 2},
		{"*1", 11},
		{"3*", 9},
		{"111", 3},
		{"226", 3},
		{"*********", 291868912},
	}

	for _, tc := range testCases {
		got := numDecodings(tc.s)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s: numDecodings(%q) = %d (want %d)\n", status, tc.s, got, tc.want)
	}
}

func numDecodings(s string) int {
	n := len(s)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// dp0 = dp[i], dp1 = dp[i-1], dp2 = dp[i-2]
	dp0, dp1, dp2 := 0, 1, 0

	for i := 0; i < n; i++ {
		dp0 = 0

		// Single digit
		if s[i] == '*' {
			dp0 = (dp0 + dp1*9) % mod
		} else if s[i] != '0' {
			dp0 = (dp0 + dp1) % mod
		}

		// Two digits
		if i > 0 {
			if s[i-1] == '*' && s[i] == '*' {
				// ** = 11-19, 21-26 = 15 possibilities
				dp0 = (dp0 + dp2*15) % mod
			} else if s[i-1] == '*' {
				// *c
				if s[i] <= '6' {
					dp0 = (dp0 + dp2*2) % mod // 1c or 2c
				} else {
					dp0 = (dp0 + dp2*1) % mod // 1c only
				}
			} else if s[i] == '*' {
				// c*
				if s[i-1] == '1' {
					dp0 = (dp0 + dp2*9) % mod // 11-19
				} else if s[i-1] == '2' {
					dp0 = (dp0 + dp2*6) % mod // 21-26
				}
			} else {
				// cc
				twoDigit := (s[i-1]-'0')*10 + (s[i] - '0')
				if twoDigit >= 10 && twoDigit <= 26 {
					dp0 = (dp0 + dp2) % mod
				}
			}
		}

		dp2 = dp1
		dp1 = dp0
	}

	return dp0
}
```
