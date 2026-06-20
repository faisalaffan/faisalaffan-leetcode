# 1387 — Sort Integers By The Power Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func getKth(lo int, hi int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP, Sorting

**Waktu:** O(n log n) for sorting  |  **Ruang:** O(n) for memoization and sorted array

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
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
	pairs := make([]pair, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		pairs = append(pairs, pair{i, power(i)})
	}

  // Custom sort
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].power != pairs[j].power {
			return pairs[i].power < pairs[j].power
		}
		return pairs[i].val < pairs[j].val
	})

	return pairs[k-1].val
}
```
