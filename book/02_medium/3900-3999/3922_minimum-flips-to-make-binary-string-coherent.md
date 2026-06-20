# 3922 — Minimum Flips To Make Binary String Coherent

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinimumFlipsToMakeBinaryStringCoherent(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N)  |  **Ruang:** O(N)


## 💻 Solusi Go

```go
package main

// LeetCode #3922: Minimum Flips to Make Binary String Coherent
// https://leetcode.com/problems/minimum-flips-to-make-binary-string-coherent/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Valid strings avoid "011" and "110" subsequences. Valid patterns:
// all zeros, all ones, exactly one 1 (0*10*), or exactly two 1s with zeros
// between (10+1). Compute min flips to each pattern.

import "fmt"

func MinimumFlipsToMakeBinaryStringCoherent(s string) int {
	n := len(s)

  // Alokasi slice
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i]
		if s[i] == '1' {
			pref[i+1]++
		}
	}
	totalOnes := pref[n]

	// Cat 1: all zeros
	ans := totalOnes

	// Cat 2: all ones
	if n-totalOnes < ans {
		ans = n - totalOnes
	}

	// Cat 3: exactly one 1 (0*10*)
	for i := 0; i < n; i++ {
		onesBefore := pref[i]
		onesAfter := pref[n] - pref[i+1]
		cost := onesBefore + onesAfter
		if s[i] == '0' {
			cost++
		}
		if cost < ans {
			ans = cost
		}
	}

	// Cat 4: 10+1 (two 1s with zeros in between)
	if n >= 2 {
		cost := 0
		if s[0] == '0' {
			cost++
		}
		if s[n-1] == '0' {
			cost++
		}
		if n > 2 {
			// Middle positions 1..n-2 must be 0
			cost += pref[n-1] - pref[1]
		}
		if cost < ans {
			ans = cost
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumFlipsToMakeBinaryStringCoherent("1010")) // Expected: 1

	// Example 2
	fmt.Println(MinimumFlipsToMakeBinaryStringCoherent("0110")) // Expected: 1

	// Example 3
	fmt.Println(MinimumFlipsToMakeBinaryStringCoherent("1000")) // Expected: 0
}
```
