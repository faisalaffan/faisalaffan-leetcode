# 2266 — Count Number Of Texts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countTexts(pressedKeys string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2266: Count Number of Texts
// https://leetcode.com/problems/count-number-of-texts/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countTexts(pressedKeys string) int {
	const mod = 1_000_000_007
	n := len(pressedKeys)
  // Alokasi slice
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] // press once
		// Check for multiple presses of same digit
		maxPress := 3
		if pressedKeys[i-1] == '7' || pressedKeys[i-1] == '9' {
			maxPress = 4
		}
		for j := 2; j <= maxPress && j <= i; j++ {
			if pressedKeys[i-j] == pressedKeys[i-1] {
				dp[i] = (dp[i] + dp[i-j]) % mod
			} else {
				break
			}
		}
	}
	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(countTexts("22233"))
	// Expected: 8

	// Test case 2
	fmt.Println(countTexts("222222222222222222222222222222222222"))
	// Expected: 82876089

	// Test case 3
	fmt.Println(countTexts("33"))
	// Expected: 2
}
```
