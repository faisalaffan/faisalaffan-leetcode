# 0461 — Hamming Distance

## Deskripsi

**Soal:** [0461. Hamming Distance](https://leetcode.com/problems/hamming-distance/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func HammingDistance(x, y int) int`

## Solusi Go

```go
package main

// LeetCode #461: Hamming Distance
// https://leetcode.com/problems/hamming-distance/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func HammingDistance(x, y int) int {
	xor := x ^ y
	count := 0
	for xor > 0 {
		xor &= xor - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(HammingDistance(1, 4))
	fmt.Println(HammingDistance(3, 1))
}
```
