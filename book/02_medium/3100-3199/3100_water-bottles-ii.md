# 3100 — Water Bottles Ii

## Deskripsi

**Soal:** [3100. Water Bottles Ii](https://leetcode.com/problems/water-bottles-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxBottlesDrunk(numBottles int, numExchange int) int`

## Solusi Go

```go
package main

// LeetCode #3100: Water Bottles II
// https://leetcode.com/problems/water-bottles-ii/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func maxBottlesDrunk(numBottles int, numExchange int) int {
	total := numBottles
	empty := numBottles

	for empty >= numExchange {
		empty -= numExchange
		numExchange++
		total++
		empty++
	}

	return total
}

func main() {
	fmt.Println(maxBottlesDrunk(13, 6)) // Expected: 15
	fmt.Println(maxBottlesDrunk(10, 3)) // Expected: 13
}
```
