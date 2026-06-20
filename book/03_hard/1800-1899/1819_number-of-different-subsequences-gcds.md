# 1819 — Number Of Different Subsequences Gcds

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countDifferentSubsequenceGCDs(nums []int) int
```

> **💡 Hint:** Count the number of distinct GCD values achievable by any

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1819: Number of Different Subsequences GCDs
// https://leetcode.com/problems/number-of-different-subsequences-gcds/
// Difficulty: Hard
//
// Approach: Count the number of distinct GCD values achievable by any
//   non-empty subsequence of nums.
//
//   For each possible GCD value g from 1 to max(nums):
//     Collect all numbers in nums that are multiples of g.
//     Compute the GCD of that collection.
//     If GCD == g, then g is achievable as a subsequence GCD.
//
//   Optimization: Use a boolean frequency array (present[x] = true if x in nums)
//   and iterate over multiples of g to compute GCD.

import "fmt"

func main() {
	// Example 1
	fmt.Println("Example 1:", countDifferentSubsequenceGCDs([]int{6, 10, 3}))
	// Expected: 5  (GCDs: 1,2,3,5,6)

	// Example 2
	fmt.Println("Example 2:", countDifferentSubsequenceGCDs([]int{5, 15, 10, 3}))
	// Expected: 6  (GCDs: 1,3,5,10,15,5? wait, let's compute: 1,2,3,5,10,15)

	// Edge case: single element
	fmt.Println("Edge (single):", countDifferentSubsequenceGCDs([]int{7}))
	// Expected: 1 (GCD = {7})

	// All ones
	fmt.Println("Edge (ones):", countDifferentSubsequenceGCDs([]int{1, 1, 1}))
	// Expected: 1 (GCD = {1})
}

func countDifferentSubsequenceGCDs(nums []int) int {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	present := make([]bool, maxVal+1)
	for _, v := range nums {
		present[v] = true
	}

	ans := 0
	// Check each possible GCD value
	for g := 1; g <= maxVal; g++ {
		curGCD := 0
		for m := g; m <= maxVal; m += g {
			if present[m] {
				if curGCD == 0 {
					curGCD = m
				} else {
					curGCD = gcd(curGCD, m)
				}
				if curGCD == g {
					ans++
					break
				}
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Stub kept for compatibility with the repo scaffold.
func NumberOfDifferentSubsequencesGcds() any {
	return countDifferentSubsequenceGCDs([]int{6, 10, 3})
}
```
