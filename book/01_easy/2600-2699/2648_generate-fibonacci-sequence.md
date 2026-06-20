# 2648 — Generate Fibonacci Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func GenerateFibonacciSequence() func() int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

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
