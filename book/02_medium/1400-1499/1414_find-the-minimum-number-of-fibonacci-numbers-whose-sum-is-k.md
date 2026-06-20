# 1414 — Find The Minimum Number Of Fibonacci Numbers Whose Sum Is K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func findMinFibonacciNumbers(k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log k) since Fibonacci numbers grow exponentially  |  **Ruang:** O(1)


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
