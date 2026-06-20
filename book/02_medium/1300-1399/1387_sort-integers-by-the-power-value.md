# 1387 — Sort Integers By The Power Value

## Deskripsi

**Soal:** [1387. Sort Integers By The Power Value](https://leetcode.com/problems/sort-integers-by-the-power-value/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(n) for memoization and sorted array

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1387: Sort Integers by The Power Value
// https://leetcode.com/problems/sort-integers-by-the-power-value/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(getKth(12, 15, 2)) // 13

	// Test case 2
	fmt.Println(getKth(1, 1, 1)) // 1

	// Test case 3
	fmt.Println(getKth(7, 11, 4)) // 7

	// Test case 4
	fmt.Println(getKth(10, 20, 5)) // 13
}

// Time: O(n log n) for sorting
// Space: O(n) for memoization and sorted array
func getKth(lo int, hi int, k int) int {
  // Membuat map untuk pencarian O(1): key → value
	memo := make(map[int]int)
	memo[1] = 0

	var power func(int) int
	power = func(x int) int {
		if val, ok := memo[x]; ok {
			return val
		}
		if x%2 == 0 {
			memo[x] = 1 + power(x/2)
		} else {
			memo[x] = 1 + power(3*x+1)
		}
		return memo[x]
	}

	type pair struct {
		val, power int
	}
  // Membuat slice untuk menyimpan hasil
	pairs := make([]pair, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		pairs = append(pairs, pair{i, power(i)})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].power != pairs[j].power {
			return pairs[i].power < pairs[j].power
		}
		return pairs[i].val < pairs[j].val
	})

	return pairs[k-1].val
}
```
