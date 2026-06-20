# 2614 — Prime In Diagonal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func isPrime(n int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * sqrt(max(nums)))  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2614: Prime In Diagonal
// https://leetcode.com/problems/prime-in-diagonal/
// Difficulty: Easy
// Time: O(n * sqrt(max(nums))) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(diagonalPrime([][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}))
	fmt.Println(diagonalPrime([][]int{{1, 2, 3}, {5, 17, 7}, {9, 11, 10}}))
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func diagonalPrime(nums [][]int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		if isPrime(nums[i][i]) && nums[i][i] > ans {
			ans = nums[i][i]
		}
		if isPrime(nums[i][n-i-1]) && nums[i][n-i-1] > ans {
			ans = nums[i][n-i-1]
		}
	}
	return ans
}
```
