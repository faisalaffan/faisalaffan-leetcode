# 1726 — Tuple With Same Product

## Deskripsi

**Soal:** [1726. Tuple With Same Product](https://leetcode.com/problems/tuple-with-same-product/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2), Space: O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** —

**Fungsi Solusi:** `func tupleSameProduct(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #1726: Tuple with Same Product
// https://leetcode.com/problems/tuple-with-same-product/
// Difficulty: Medium
// Time: O(n^2), Space: O(n^2)

import "fmt"

func tupleSameProduct(nums []int) int {
	n := len(nums)
  // Membuat map untuk pencarian O(1): key → value
	productCount := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			product := nums[i] * nums[j]
			productCount[product]++
		}
	}

	result := 0
	for _, count := range productCount {
		if count > 1 {
			// Each pair of pairs = 8 tuples (4! / 3 = 8)
			result += count * (count - 1) / 2 * 8
		}
	}
	return result
}

func main() {
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6}))     // Expected: 8
	fmt.Println(tupleSameProduct([]int{1, 2, 4, 5, 10})) // Expected: 16
	fmt.Println(tupleSameProduct([]int{1, 2, 3, 4, 6, 12})) // Expected: 40
}
```
