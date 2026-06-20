# 0201 — Bitwise And Of Numbers Range

## Deskripsi

**Soal:** [0201. Bitwise And Of Numbers Range](https://leetcode.com/problems/bitwise-and-of-numbers-range/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log max(left, right)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func rangeBitwiseAnd(left int, right int) int`

## Solusi Go

```go
package main

// LeetCode #201: Bitwise AND of Numbers Range
// https://leetcode.com/problems/bitwise-and-of-numbers-range/
// Difficulty: Medium
// Time: O(log max(left, right)), Space: O(1)

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
	shift := 0
  // Loop two-pointer: kiri vs kanan
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 0))
	fmt.Println(rangeBitwiseAnd(1, 2147483647))
}
```
