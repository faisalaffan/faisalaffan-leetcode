# 2466 — Count Ways To Build Good Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func countGoodStrings(low int, high int, zero int, one int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(high)  
**Kompleksitas Ruang:** O(high)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2466: Count Ways To Build Good Strings
// https://leetcode.com/problems/count-ways-to-build-good-strings/
// Difficulty: Medium
// Time: O(high) | Space: O(high)
// DP: dp[i] = ways to build string of length i.
// dp[i] = dp[i-zero] + dp[i-one] (if i >= zero/one).

import "fmt"

func main() {
	fmt.Println(countGoodStrings(3, 3, 1, 1)) // 8
	fmt.Println(countGoodStrings(2, 3, 1, 2)) // 5
}

const MOD = 1000000007

func countGoodStrings(low int, high int, zero int, one int) int {
  // Alokasi slice integer
	dp := make([]int, high+1)
	dp[0] = 1
	ans := 0

	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % MOD
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % MOD
		}
		if i >= low {
			ans = (ans + dp[i]) % MOD
		}
	}
	return ans
}
```
