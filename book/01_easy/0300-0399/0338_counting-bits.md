# 0338 — Counting Bits

## Deskripsi

**Soal:** [0338. Counting Bits](https://leetcode.com/problems/counting-bits/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func CountingBits(n int) []int`

## Solusi Go

```go
package main

// LeetCode #338: Counting Bits
// https://leetcode.com/problems/counting-bits/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func CountingBits(n int) []int {
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return ans
}

func main() {
	fmt.Println(CountingBits(2))
	fmt.Println(CountingBits(5))
	fmt.Println(CountingBits(0))
}
```
