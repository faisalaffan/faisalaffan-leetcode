# 1414 — Find The Minimum Number Of Fibonacci Numbers Whose Sum Is K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func findMinFibonacciNumbers(k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log k) since Fibonacci numbers grow exponentially  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1414: Find the Minimum Number of Fibonacci Numbers Whose Sum Is K
// https://leetcode.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(findMinFibonacciNumbers(7)) // 2

	// Test case 2
	fmt.Println(findMinFibonacciNumbers(10)) // 2

	// Test case 3
	fmt.Println(findMinFibonacciNumbers(19)) // 3

	// Test case 4
	fmt.Println(findMinFibonacciNumbers(1)) // 1
}

// Time: O(log k) since Fibonacci numbers grow exponentially
// Space: O(1)
func findMinFibonacciNumbers(k int) int {
	// Generate all Fibonacci numbers <= k
	fib := []int{1, 1}
	for fib[len(fib)-1] <= k {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		fib = append(fib, next)
	}

	count := 0
	remaining := k
	for i := len(fib) - 1; i >= 0; i-- {
		if fib[i] <= remaining {
			remaining -= fib[i]
			count++
		}
		if remaining == 0 {
			break
		}
	}

	return count
}
```
