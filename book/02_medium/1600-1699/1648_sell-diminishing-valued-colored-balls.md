# 1648 — Sell Diminishing Valued Colored Balls

## Deskripsi

**Soal:** [1648. Sell Diminishing Valued Colored Balls](https://leetcode.com/problems/sell-diminishing-valued-colored-balls/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N log N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1648: Sell Diminishing-Valued Colored Balls
// https://leetcode.com/problems/sell-diminishing-valued-colored-balls/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxProfit([]int{2, 5}, 4))
	fmt.Println(MaxProfit([]int{3, 5}, 6))
	fmt.Println(MaxProfit([]int{2, 8, 4, 10, 6}, 20))
}

func MaxProfit(inventory []int, orders int) int {
	// Time: O(N log N), Space: O(N)
	const mod = 1_000_000_007

	// Sort descending
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i] > inventory[j]
	})

	// Append 0 for convenience
	inventory = append(inventory, 0)
	n := len(inventory)

	profit := 0
	count := 0

	for i := 0; i < n-1; i++ {
		if inventory[i] == inventory[i+1] {
			continue
		}

		// Height difference between current and next level
		height := inventory[i] - inventory[i+1]
		width := i + 1 // number of types with this count
		total := height * width

		if count+total <= orders {
			// Take all balls at this level
			// Sum for this level: width * sum of (inventory[i], inventory[i]-1, ..., inventory[i+1]+1)
			top := inventory[i]
			bottom := inventory[i+1] + 1
			sumLevel := (top + bottom) * height / 2
			profit = (profit + sumLevel*width) % mod
			count += total
		} else {
			// Take only some
			remaining := orders - count
			fullRows := remaining / width
			extra := remaining % width

			top := inventory[i]
			bottom := top - fullRows + 1
			sumFull := (top + bottom) * fullRows / 2
			profit = (profit + sumFull*width) % mod
			profit = (profit + bottom*extra) % mod

			count = orders
			break
		}
	}

	return profit
}
```
