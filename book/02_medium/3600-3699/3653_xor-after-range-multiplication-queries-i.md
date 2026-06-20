# 3653 — Xor After Range Multiplication Queries I

## Deskripsi

**Soal:** [3653. Xor After Range Multiplication Queries I](https://leetcode.com/problems/xor-after-range-multiplication-queries-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func xorAfterRangeMultiplicationQueriesI(arr []int, queries [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #3653: XOR After Range Multiplication Queries I
// https://leetcode.com/problems/xor-after-range-multiplication-queries-i/
// Difficulty: Medium
// Time: O(n + q) | Space: O(1)

import "fmt"

func xorAfterRangeMultiplicationQueriesI(arr []int, queries [][]int) []int {
	n := len(arr)
  // Membuat slice untuk menyimpan hasil
	pref := make([]int, n+1)
	for i, v := range arr {
		pref[i+1] = pref[i] ^ v
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r, mul := q[0], q[1], q[2]
		// Apply multiplication to the range [l, r]
		for i := l; i <= r; i++ {
			arr[i] *= mul
		}
		// Recompute prefix XOR
		for i := l; i <= r; i++ {
			pref[i+1] = pref[i] ^ arr[i]
		}
		// XOR of range [l, r]
		ans[qi] = pref[r+1] ^ pref[l]
	}
	return ans
}

func main() {
	fmt.Println(xorAfterRangeMultiplicationQueriesI([]int{1, 2, 3}, [][]int{{0, 1, 2}, {1, 2, 3}}))
	fmt.Println(xorAfterRangeMultiplicationQueriesI([]int{5, 7}, [][]int{{0, 0, 2}, {0, 1, 1}}))
}
```
