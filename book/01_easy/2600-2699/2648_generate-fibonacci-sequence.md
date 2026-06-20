# 2648 — Generate Fibonacci Sequence

## Deskripsi

**Soal:** [2648. Generate Fibonacci Sequence](https://leetcode.com/problems/generate-fibonacci-sequence/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2648: Generate Fibonacci Sequence
// https://leetcode.com/problems/generate-fibonacci-sequence/
// Difficulty: Easy
// Time: O(n) | Space: O(1)
// Note: JavaScript generator problem, adapted to Go.

import "fmt"

func main() {
	fib := GenerateFibonacciSequence()
	for i := 0; i < 5; i++ {
		fmt.Println(fib())
	}
}

func GenerateFibonacciSequence() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}
```
