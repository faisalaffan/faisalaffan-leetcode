# 3509 — Maximum Product Of Subsequences With An Alternating Sum Equal To K

## Deskripsi

**Soal:** [3509. Maximum Product Of Subsequences With An Alternating Sum Equal To K](https://leetcode.com/problems/maximum-product-of-subsequences-with-an-alternating-sum-equal-to-k/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** DP with maps. Track (sum, parity) -> max product, with products

## Solusi Go

```go
package main

// LeetCode #3509: Maximum Product of Subsequences With an Alternating Sum Equal to K
// https://leetcode.com/problems/maximum-product-of-subsequences-with-an-alternating-sum-equal-to-k/
// Difficulty: Hard
//
// Find a non-empty subsequence where alternating sum (even-indexed elements
// minus odd-indexed elements) equals k. Maximize the product subject to
// product <= limit. Return max product, or -1 if none.
//
// Approach: DP with maps. Track (sum, parity) -> max product, with products
// capped at limit+1.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxProduct([]int{1, 2, 3}, 2, 10))
	// Example 2
	fmt.Println(maxProduct([]int{1, 2, 3}, 3, 5))
	// Example 3
	fmt.Println(maxProduct([]int{0, 1, 2}, 1, 100))
	// Edge: single element
	fmt.Println(maxProduct([]int{5}, 5, 100))
	// Edge: no valid subsequence
	fmt.Println(maxProduct([]int{1, 1}, 5, 10))
}

func maxProduct(nums []int, k int, limit int) int {
	// even[sum] = set of possible products for subsequences with even length (next sign is +)
	// odd[sum] = set of possible products for subsequences with odd length (next sign is -)
  // Membuat map untuk pencarian O(1): key → value
	even := make(map[int]map[int]bool)
  // Membuat map untuk pencarian O(1): key → value
	odd := make(map[int]map[int]bool)

	ans := -1

	for _, num := range nums {
		// Create new maps for current iteration
		newEven := copyMap(even)
		newOdd := copyMap(odd)

		// Start new subsequence with this element
		if newOdd[num] == nil {
			newOdd[num] = make(map[int]bool)
		}
		prod := num
		if prod > limit+1 {
			prod = limit + 1
		}
		newOdd[num][prod] = true
		if num == k && prod <= limit {
			if prod > ans {
				ans = prod
			}
		}

		// Extend existing subsequences
		// Extend even-length: add num (odd length becomes)
		for sum, prods := range even {
			newSum := sum + num
			for p := range prods {
				newProd := p * num
				if newProd > limit+1 {
					newProd = limit + 1
				}
				if newOdd[newSum] == nil {
					newOdd[newSum] = make(map[int]bool)
				}
				newOdd[newSum][newProd] = true
				if newSum == k && newProd <= limit && newProd > ans {
					ans = newProd
				}
			}
		}

		// Extend odd-length: subtract num (even length becomes)
		for sum, prods := range odd {
			newSum := sum - num
			for p := range prods {
				newProd := p * num
				if newProd > limit+1 {
					newProd = limit + 1
				}
				if newEven[newSum] == nil {
					newEven[newSum] = make(map[int]bool)
				}
				newEven[newSum][newProd] = true
				if newSum == k && newProd <= limit && newProd > ans {
					ans = newProd
				}
			}
		}

		even = newEven
		odd = newOdd
	}

	return ans
}

func copyMap(src map[int]map[int]bool) map[int]map[int]bool {
  // Membuat map untuk pencarian O(1): key → value
	dst := make(map[int]map[int]bool)
	for k, v := range src {
		dst[k] = make(map[int]bool)
		for p := range v {
			dst[k][p] = true
		}
	}
	return dst
}
```
