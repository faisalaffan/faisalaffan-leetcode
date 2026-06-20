# 0974 — Subarray Sums Divisible By K

## Deskripsi

**Soal:** [0974. Subarray Sums Divisible By K](https://leetcode.com/problems/subarray-sums-divisible-by-k/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(k)

**Algoritma:** Prefix Sum (jumlah kumulatif)

> **Ide Kunci:** Prefix sum + modulo counting

## Solusi Go

```go
package main

// LeetCode #974: Subarray Sums Divisible by K
// https://leetcode.com/problems/subarray-sums-divisible-by-k/
// Difficulty: Medium
//
// Approach: Prefix sum + modulo counting
// Time: O(n)
// Space: O(k)

import "fmt"

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)) // 7
	fmt.Println(subarraysDivByK([]int{5}, 9))                  // 0
	fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))           // 2
}

func subarraysDivByK(nums []int, k int) int {
  // Membuat map untuk pencarian O(1): key → value
	modCount := make(map[int]int)
	modCount[0] = 1
	prefixSum := 0
	result := 0

	for _, n := range nums {
		prefixSum += n
		mod := prefixSum % k
		if mod < 0 {
			mod += k
		}
		result += modCount[mod]
		modCount[mod]++
	}

	return result
}
```
