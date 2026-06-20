# 1655 — Distribute Repeating Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func canDistribute(nums []int, quantity []int) bool
```

> **💡 Hint:** // 1. Count frequencies of each distinct number in nums.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming, Bitmask

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
	"sort"
)

// LeetCode #1655: Distribute Repeating Integers
// https://leetcode.com/problems/distribute-repeating-integers/
// Difficulty: Hard
//
// We have an array nums of integers (may have duplicates) and an array quantity
// where quantity[j] is the number of items that customer j wants to order.
// Each customer must receive items of the SAME value (all identical integers).
// Determine if it's possible to satisfy all customers.
//
// Approach:
// 1. Count frequencies of each distinct number in nums.
// 2. Sort frequencies descending (larger counts first) for pruning.
// 3. Sort quantity descending (larger orders first) for better pruning.
// 4. Use DP with bitmask: can[mask] = true if the current set of frequencies
//    can satisfy the subset of customers represented by mask.
// 5. Iterate over frequencies; for each, update the DP.

func canDistribute(nums []int, quantity []int) bool {
	// Count frequencies
  // Membuat map (HashMap) — pencarian O(1)
	freqMap := make(map[int]int)
	for _, v := range nums {
		freqMap[v]++
	}

  // Alokasi slice integer
	freqs := make([]int, 0, len(freqMap))
	for _, f := range freqMap {
		freqs = append(freqs, f)
	}

	// Sort frequencies descending for better pruning
  // Custom sort dengan comparator
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i] > freqs[j]
	})

	// Sort quantity descending
  // Custom sort dengan comparator
	sort.Slice(quantity, func(i, j int) bool {
		return quantity[i] > quantity[j]
	})

	n := len(quantity)
	size := 1 << n

	// Precompute subset sums of quantity
  // Alokasi slice integer
	subsetSum := make([]int, size)
	for mask := 1; mask < size; mask++ {
		lsb := mask & -mask
		bit := int(math.Log2(float64(lsb)))
		subsetSum[mask] = subsetSum[mask^lsb] + quantity[bit]
	}

	// dp[mask] = true if we can satisfy customer subset 'mask' with processed frequencies
	dp := make([]bool, size)
	dp[0] = true

	for _, freq := range freqs {
		// For each mask that is currently achievable, try adding this frequency
		// to cover additional customers.
		// We need to iterate backwards to avoid using the same frequency multiple times.
		for mask := size - 1; mask >= 0; mask-- {
			if !dp[mask] {
				continue
			}
			// Find the complement (customers not yet served)
			remaining := (size - 1) ^ mask
			// Try all subsets of remaining customers
			sub := remaining
			for sub > 0 {
				if subsetSum[sub] <= freq {
					dp[mask|sub] = true
				}
				sub = (sub - 1) & remaining
			}
		}

		if dp[size-1] {
			return true
		}
	}

	return dp[size-1]
}

func main() {
	// Example 1:
	// Input: nums = [1,2,3,4], quantity = [2]
	// Output: false — we have 4 distinct numbers with frequency 1 each,
	// but customer needs 2 items of the same value.
	fmt.Println(canDistribute([]int{1, 2, 3, 4}, []int{2}))

	// Example 2:
	// Input: nums = [1,2,3,3], quantity = [2]
	// Output: true — we have two 3's, give them both to the customer.
	fmt.Println(canDistribute([]int{1, 2, 3, 3}, []int{2}))

	// Example 3:
	// Input: nums = [1,1,2,2], quantity = [2,2]
	// Output: true — give 1,1 to one customer, 2,2 to the other.
	fmt.Println(canDistribute([]int{1, 1, 2, 2}, []int{2, 2}))

	// Example 4:
	// Input: nums = [1,1,2,3], quantity = [2,2]
	// Output: false — only one pair (1,1), need two pairs.
	fmt.Println(canDistribute([]int{1, 1, 2, 3}, []int{2, 2}))

	// Example 5:
	// Input: nums = [1,1,1,1,1], quantity = [2,3]
	// Output: true — give 2 of the 1's to one customer, 3 to the other.
	fmt.Println(canDistribute([]int{1, 1, 1, 1, 1}, []int{2, 3}))
}
```
