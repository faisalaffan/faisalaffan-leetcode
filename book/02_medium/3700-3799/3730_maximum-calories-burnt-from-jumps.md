# 3730 — Maximum Calories Burnt From Jumps

## Deskripsi

**Soal:** [3730. Maximum Calories Burnt From Jumps](https://leetcode.com/problems/maximum-calories-burnt-from-jumps/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func maximumCaloriesBurntFromJumps(heights []int) int64`

## Solusi Go

```go
package main

// LeetCode #3730: Maximum Calories Burnt from Jumps
// https://leetcode.com/problems/maximum-calories-burnt-from-jumps/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumCaloriesBurntFromJumps(heights []int) int64 {
	n := len(heights)
	sort.Ints(heights)

  // Membuat slice untuk menyimpan hasil
	seq := make([]int, 0, n)
	l, r := 0, n-1
	for l <= r {
		seq = append(seq, heights[r])
		r--
		if l <= r {
			seq = append(seq, heights[l])
			l++
		}
	}

	total := int64(seq[0]) * int64(seq[0])
	for i := 1; i < n; i++ {
		diff := seq[i] - seq[i-1]
		total += int64(diff) * int64(diff)
	}
	return total
}

func main() {
	fmt.Println(maximumCaloriesBurntFromJumps([]int{1, 7, 9}))
	fmt.Println(maximumCaloriesBurntFromJumps([]int{5, 2, 4}))
	fmt.Println(maximumCaloriesBurntFromJumps([]int{3, 3}))
}
```
