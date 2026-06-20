# 1416 — Restore The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func numberOfArrays(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1416: Restore The Array
// https://leetcode.com/problems/restore-the-array/
// Difficulty: Hard

import "fmt"

const mod1416 = 1_000_000_007

func numberOfArrays(s string, k int) int {
	n := len(s)
  // Alokasi slice
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}
		var num int
		for j := i; j < n; j++ {
			num = num*10 + int(s[j]-'0')
			if num > k {
				break
			}
			dp[i] = (dp[i] + dp[j+1]) % mod1416
		}
	}
	return dp[0]
}

func main() {
	// Example: "1317", 2000 -> 8
	fmt.Println(numberOfArrays("1317", 2000))
}
```
