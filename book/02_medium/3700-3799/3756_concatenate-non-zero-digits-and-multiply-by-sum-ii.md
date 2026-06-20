# 3756 — Concatenate Non Zero Digits And Multiply By Sum Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func concatenateNonZeroDigitsAndMultiplyBySumIi(s string, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n + q)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3756: Concatenate Non-Zero Digits and Multiply by Sum II
// https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-ii/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

const mod3756 = 1000000007

func concatenateNonZeroDigitsAndMultiplyBySumIi(s string, queries [][]int) []int {
	n := len(s)
  // Alokasi slice
	prefixCnt := make([]int, n+1)
  // Alokasi slice
	prefixSum := make([]int, n+1)
  // Alokasi slice
	prefixNum := make([]int64, n+1)

	for i, ch := range s {
		d := int(ch - '0')
		prefixSum[i+1] = prefixSum[i] + d
		if d > 0 {
			prefixCnt[i+1] = prefixCnt[i] + 1
			prefixNum[i+1] = (prefixNum[i]*10 + int64(d)) % mod3756
		} else {
			prefixCnt[i+1] = prefixCnt[i]
			prefixNum[i+1] = prefixNum[i]
		}
	}

  // Alokasi slice
	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		subLen := prefixCnt[r+1] - prefixCnt[l]
		// x = (prefixNum[r+1] - prefixNum[l] * 10^subLen) % mod
		x := (prefixNum[r+1] - prefixNum[l]*pow10(int64(subLen))) % mod3756
		if x < 0 {
			x += mod3756
		}
		digitSum := prefixSum[r+1] - prefixSum[l]
		ans[qi] = int((x * int64(digitSum)) % mod3756)
	}
	return ans
}

func pow10(exp int64) int64 {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return 10
	}
	half := pow10(exp / 2)
	half = (half * half) % mod3756
	if exp%2 == 1 {
		half = (half * 10) % mod3756
	}
	return half
}

func main() {
	fmt.Println(concatenateNonZeroDigitsAndMultiplyBySumIi("10203004", [][]int{{0, 7}, {1, 3}, {4, 6}}))
	fmt.Println(concatenateNonZeroDigitsAndMultiplyBySumIi("1000", [][]int{{0, 3}, {1, 1}}))
	fmt.Println(concatenateNonZeroDigitsAndMultiplyBySumIi("9876543210", [][]int{{0, 9}}))
}
```
