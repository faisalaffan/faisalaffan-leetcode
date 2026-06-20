# 1774 — Closest Dessert Cost

## Deskripsi

**Soal:** [1774. Closest Dessert Cost](https://leetcode.com/problems/closest-dessert-cost/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(b * 3^t), Space: O(t)  
**Kompleksitas Ruang:** O(t)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func closestCost(baseCosts []int, toppingCosts []int, target int) int`

## Solusi Go

```go
package main

// LeetCode #1774: Closest Dessert Cost
// https://leetcode.com/problems/closest-dessert-cost/
// Difficulty: Medium
// Time: O(b * 3^t), Space: O(t)

import "fmt"

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	best := baseCosts[0]
	for _, b := range baseCosts {
		best = minDiff(best, b, target)
		dfs(toppingCosts, 0, b, target, &best)
	}
	return best
}

func dfs(toppings []int, idx int, current int, target int, best *int) {
	if idx == len(toppings) {
		*best = minDiff(*best, current, target)
		return
	}
	// Try 0, 1, or 2 of each topping
	for count := 0; count <= 2; count++ {
		dfs(toppings, idx+1, current+count*toppings[idx], target, best)
	}
}

func minDiff(a, b, target int) int {
	diffA := abs(a - target)
	diffB := abs(b - target)
	if diffA < diffB || (diffA == diffB && a < b) {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(closestCost([]int{1, 7}, []int{3, 4}, 10)) // Expected: 10
	fmt.Println(closestCost([]int{2, 3}, []int{4, 5, 100}, 18)) // Expected: 17
	fmt.Println(closestCost([]int{10}, []int{1}, 1)) // Expected: 10
}
```
