# 3610 — Minimum Number Of Primes To Sum To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MinimumNumberOfPrimesToSumToTarget(target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3610: Minimum Number of Primes to Sum to Target
// https://leetcode.com/problems/minimum-number-of-primes-to-sum-to-target/
// Difficulty: Medium [Paid]
// Complexity: O(target * log log target) time, O(target) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumNumberOfPrimesToSumToTarget(10))
	// Test case 2
	fmt.Println("Test 2:", MinimumNumberOfPrimesToSumToTarget(3))
	// Test case 3
	fmt.Println("Test 3:", MinimumNumberOfPrimesToSumToTarget(1))
}

func MinimumNumberOfPrimesToSumToTarget(target int) int {
	if target < 2 {
		return -1
	}
	// Sieve to find all primes up to target
	isPrime := make([]bool, target+1)
	for i := 2; i <= target; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= target; i++ {
		if isPrime[i] {
			for j := i * i; j <= target; j += i {
				isPrime[j] = false
			}
		}
	}

	// DP: min primes to sum to x
  // Alokasi slice
	dp := make([]int, target+1)
  // Range loop
	for i := range dp {
		dp[i] = target + 1
	}
	dp[0] = 0
	for i := 2; i <= target; i++ {
		if isPrime[i] {
			for j := i; j <= target; j++ {
				if dp[j-i]+1 < dp[j] {
					dp[j] = dp[j-i] + 1
				}
			}
		}
	}
	if dp[target] > target {
		return -1
	}
	return dp[target]
}
```
