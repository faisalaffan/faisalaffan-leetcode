# 2518 — Number Of Great Partitions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfGreatPartitions(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2518: Number of Great Partitions
// https://leetcode.com/problems/number-of-great-partitions/
// Difficulty: Hard

import "fmt"

const mod = 1000000007

// numberOfGreatPartitions counts the number of ways to partition nums into
// two groups such that |sum(group1) - sum(group2)| >= k.
//
// Total ways = 2^n. Subtract ways where |2*subsetSum - totalSum| < k.
// DP counts subsets achieving each possible sum.
//
// Complexity: O(n * totalSum) time, O(totalSum) space
func numberOfGreatPartitions(nums []int, k int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	// dp[s] = number of ways to achieve subset sum s (using distinct elements)
  // Alokasi slice integer
	dp := make([]int, totalSum+1)
	dp[0] = 1
	for _, v := range nums {
		for s := totalSum; s >= v; s-- {
			dp[s] = (dp[s] + dp[s-v]) % mod
		}
	}

	// Count "bad" partitions where |2*sum - totalSum| < k
	bad := 0
	for sum := 0; sum <= totalSum; sum++ {
		if abs(2*sum-totalSum) < k {
			bad = (bad + dp[sum]) % mod
		}
	}

	// Total partitions = 2^n
	total := 1
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		total = (total * 2) % mod
	}

	return (total - bad + mod) % mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: nums=[1,2,3,4], k=4 ->", numberOfGreatPartitions([]int{1, 2, 3, 4}, 4))

	// Additional test cases
	fmt.Println("Test 2: nums=[1,1,1,1], k=2 ->", numberOfGreatPartitions([]int{1, 1, 1, 1}, 2))
	fmt.Println("Test 3: nums=[1,2], k=1 ->", numberOfGreatPartitions([]int{1, 2}, 1))
	fmt.Println("Test 4: nums=[1], k=1 ->", numberOfGreatPartitions([]int{1}, 1))
	fmt.Println("Test 5: nums=[3,5], k=1 ->", numberOfGreatPartitions([]int{3, 5}, 1))
	fmt.Println("Test 6: nums=[1,2,3], k=100 ->", numberOfGreatPartitions([]int{1, 2, 3}, 100))
}
```
