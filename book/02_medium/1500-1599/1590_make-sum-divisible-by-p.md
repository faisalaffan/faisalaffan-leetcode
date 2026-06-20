# 1590 — Make Sum Divisible By P

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinSubarray(nums []int, p int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** O(N), Space: O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1590: Make Sum Divisible by P
// https://leetcode.com/problems/make-sum-divisible-by-p/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSubarray([]int{3, 1, 4, 2}, 6))
	fmt.Println(MinSubarray([]int{6, 3, 5, 2}, 9))
	fmt.Println(MinSubarray([]int{1, 2, 3}, 7))
}

func MinSubarray(nums []int, p int) int {
	// Time: O(N), Space: O(N)
	n := len(nums)

	// Total sum modulo p
	totalSum := 0
	for _, num := range nums {
		totalSum = (totalSum + num) % p
	}

	target := totalSum // the remainder we need to remove
	if target == 0 {
		return 0
	}

	// Map from prefix sum modulo p to index
  // HashMap: O(1) lookup
	prefixMap := make(map[int]int)
	prefixMap[0] = -1
	prefixSum := 0
	minLen := n

	for i, num := range nums {
		prefixSum = (prefixSum + num) % p
		// We need prefixSum - prefixSum[j] ≡ target (mod p)
		// => prefixSum[j] ≡ prefixSum - target (mod p)
		needed := (prefixSum - target + p) % p
		if j, exists := prefixMap[needed]; exists {
			if i-j < minLen {
				minLen = i - j
			}
		}
		prefixMap[prefixSum] = i
	}

	if minLen == n {
		return -1
	}
	return minLen
}
```
