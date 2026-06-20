# 2044 — Count Number Of Maximum Bitwise Or Subsets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countMaxOrSubsets(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(2^n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2044: Count Number of Maximum Bitwise-OR Subsets
// https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/
// Difficulty: Medium
// Time: O(2^n) | Space: O(n)

import "fmt"

func countMaxOrSubsets(nums []int) int {
	maxOr := 0
	for _, v := range nums {
		maxOr |= v
	}

	count := 0
	var backtrack func(idx int, curOr int)
	backtrack = func(idx int, curOr int) {
		if idx == len(nums) {
			if curOr == maxOr {
				count++
			}
			return
		}
		// Skip current
		backtrack(idx+1, curOr)
		// Take current
		backtrack(idx+1, curOr|nums[idx])
	}

	backtrack(0, 0)
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countMaxOrSubsets([]int{3, 1}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", countMaxOrSubsets([]int{2, 2, 2}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", countMaxOrSubsets([]int{3, 2, 1, 5}))
	// Expected: 6
}
```
