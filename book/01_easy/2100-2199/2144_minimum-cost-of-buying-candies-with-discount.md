# 2144 — Minimum Cost Of Buying Candies With Discount

## Deskripsi

**Soal:** [2144. Minimum Cost Of Buying Candies With Discount](https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(1) ignoring sort  
**Kompleksitas Ruang:** O(1) ignoring sort

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2144: Minimum Cost of Buying Candies With Discount
// https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumCostOfBuyingCandiesWithDiscount([]int{1, 2, 3}))    // 5
	fmt.Println(MinimumCostOfBuyingCandiesWithDiscount([]int{6, 5, 7, 9, 2, 2})) // 23
	fmt.Println(MinimumCostOfBuyingCandiesWithDiscount([]int{5, 5}))        // 10
}

// Time: O(n log n), Space: O(1) ignoring sort
func MinimumCostOfBuyingCandiesWithDiscount(cost []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(cost)))
	total := 0
	for i, c := range cost {
		if i%3 != 2 { // buy 2, get 1 free (the cheapest = every 3rd item)
			total += c
		}
	}
	return total
}
```
