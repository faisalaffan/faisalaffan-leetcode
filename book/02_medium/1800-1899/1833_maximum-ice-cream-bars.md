# 1833 — Maximum Ice Cream Bars

## Deskripsi

**Soal:** [1833. Maximum Ice Cream Bars](https://leetcode.com/problems/maximum-ice-cream-bars/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxIceCream(costs []int, coins int) int`

## Solusi Go

```go
package main

// LeetCode #1833: Maximum Ice Cream Bars
// https://leetcode.com/problems/maximum-ice-cream-bars/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	count := 0
	for _, c := range costs {
		if coins >= c {
			coins -= c
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 7)) // Expected: 4
	fmt.Println(maxIceCream([]int{10, 6, 8, 7, 7, 8}, 5)) // Expected: 0
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20)) // Expected: 6
}
```
