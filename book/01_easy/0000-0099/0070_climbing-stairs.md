# 0070 — Climbing Stairs

## Deskripsi

**Soal:** [0070. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ClimbStairs(n int) int`

## Solusi Go

```go
package main

// LeetCode #70: Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(ClimbStairs(2))
	fmt.Println(ClimbStairs(3))
	fmt.Println(ClimbStairs(4))
}
```
