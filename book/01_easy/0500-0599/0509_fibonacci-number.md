# 0509 — Fibonacci Number

## Deskripsi

**Soal:** [0509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func FibonacciNumber(n int) int`

## Solusi Go

```go
package main

// LeetCode #509: Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FibonacciNumber(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(FibonacciNumber(2))
	fmt.Println(FibonacciNumber(3))
	fmt.Println(FibonacciNumber(4))
}
```
