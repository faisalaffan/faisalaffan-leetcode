# 3618 — Split Array By Prime Indices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func isPrime(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3618: Split Array by Prime Indices
// https://leetcode.com/problems/split-array-by-prime-indices/
// Difficulty: Medium
// Complexity: O(n * sqrt(m)) time, O(1) space

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", SplitArrayByPrimeIndices([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", SplitArrayByPrimeIndices([]int{10, 20, 30, 40}))
	// Test case 3
	fmt.Println("Test 3:", SplitArrayByPrimeIndices([]int{1, 2}))
}

func SplitArrayByPrimeIndices(nums []int) [][]int {
	var result [][]int
	var current []int
	for i, v := range nums {
		current = append(current, v)
		if isPrime(i) || i == len(nums)-1 {
			result = append(result, current)
			current = nil
		}
	}
	return result
}
```
