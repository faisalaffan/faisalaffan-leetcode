# 3610 — Minimum Number Of Primes To Sum To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberOfPrimesToSumToTarget(target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Alokasi slice integer
	dp := make([]int, target+1)
  // Range loop: iterasi dengan indeks + nilai
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
