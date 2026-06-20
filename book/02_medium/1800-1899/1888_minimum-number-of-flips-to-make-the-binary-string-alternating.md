# 1888 — Minimum Number Of Flips To Make The Binary String Alternating

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinFlips(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1888: Minimum Number of Flips to Make the Binary String Alternating
// https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinFlips("111000"))
	fmt.Println(MinFlips("010"))
	fmt.Println(MinFlips("1110"))
}

// Time: O(n), Space: O(n)
func MinFlips(s string) int {
	n := len(s)
	double := s + s
	ans := n

	// Compare against "01" pattern
	mismatch0, mismatch1 := 0, 0
  // Linear scan O(n)
	for i := 0; i < len(double); i++ {
		expected0 := byte('0')
		if i%2 == 1 {
			expected0 = '1'
		}
		expected1 := byte('1')
		if i%2 == 1 {
			expected1 = '0'
		}

		if double[i] != expected0 {
			mismatch0++
		}
		if double[i] != expected1 {
			mismatch1++
		}

		if i >= n {
			left := i - n
			if double[left] != byte('0')+byte((left)%2) {
				mismatch0--
			}
			if double[left] != byte('1')-byte((left)%2) {
				mismatch1--
			}
		}

		if i >= n-1 {
			if mismatch0 < ans {
				ans = mismatch0
			}
			if mismatch1 < ans {
				ans = mismatch1
			}
		}
	}
	return ans
}
```
