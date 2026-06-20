# 3230 — Customer Purchasing Behavior Analysis

## Deskripsi

**Soal:** [3230. Customer Purchasing Behavior Analysis](https://leetcode.com/problems/customer-purchasing-behavior-analysis/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func customerPurchasingBehavior(purchases [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #3230: Customer Purchasing Behavior Analysis
// https://leetcode.com/problems/customer-purchasing-behavior-analysis/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func customerPurchasingBehavior(purchases [][]int) []int {
  // Membuat map untuk pencarian O(1): key → value
	counts := make(map[int]int)
	for _, p := range purchases {
		counts[p[0]]++
	}

	var customers []int
	for id := range counts {
		customers = append(customers, id)
	}
	sort.Slice(customers, func(i, j int) bool {
		if counts[customers[i]] != counts[customers[j]] {
			return counts[customers[i]] > counts[customers[j]]
		}
		return customers[i] < customers[j]
	})

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(customers))
	for i, id := range customers {
		ans[i] = id
	}
	return ans
}

func main() {
	fmt.Println(customerPurchasingBehavior([][]int{{1, 100}, {2, 50}, {1, 200}, {3, 75}})) // Expected: [1 2 3]
	fmt.Println(customerPurchasingBehavior([][]int{{1, 10}, {2, 20}, {2, 30}}))             // Expected: [2 1]
}
```
