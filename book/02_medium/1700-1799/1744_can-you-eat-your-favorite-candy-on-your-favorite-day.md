# 1744 — Can You Eat Your Favorite Candy On Your Favorite Day

## Deskripsi

**Soal:** [1744. Can You Eat Your Favorite Candy On Your Favorite Day](https://leetcode.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + q), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func canEat(candiesCount []int, queries [][]int) []bool`

## Solusi Go

```go
package main

// LeetCode #1744: Can You Eat Your Favorite Candy on Your Favorite Day?
// https://leetcode.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/
// Difficulty: Medium
// Time: O(n + q), Space: O(n)

import "fmt"

func canEat(candiesCount []int, queries [][]int) []bool {
	n := len(candiesCount)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + candiesCount[i]
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]bool, len(queries))
	for i, q := range queries {
		favType, favDay, dailyCap := q[0], q[1], q[2]

		// Earliest day we can eat favType candy (at dailyCap per day)
		minDay := prefix[favType] / dailyCap
		// Latest day we can eat favType candy (at 1 per day)
		maxDay := prefix[favType+1] - 1

		result[i] = favDay >= minDay && favDay <= maxDay
	}
	return result
}

func main() {
	fmt.Println(canEat([]int{7, 4, 5, 3, 8}, [][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 100}}))
	// Expected: [true, false, true]

	fmt.Println(canEat([]int{5, 2, 6, 4, 1}, [][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {0, 5, 1}}))
	// Expected: [false, true, true, false]

	fmt.Println(canEat([]int{16, 38, 8, 41, 30, 31, 14, 45, 3, 2, 24, 23, 38, 30, 4, 43}, [][]int{{8, 19, 38}}))
}
```
