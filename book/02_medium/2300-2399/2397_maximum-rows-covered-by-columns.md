# 2397 — Maximum Rows Covered By Columns

## Deskripsi

**Soal:** [2397. Maximum Rows Covered By Columns](https://leetcode.com/problems/maximum-rows-covered-by-columns/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(C(cols, select) * rows)  
**Kompleksitas Ruang:** O(cols)

**Algoritma:** Bitmask (representasi himpunan dengan bit)

## Solusi Go

```go
package main

// LeetCode #2397: Maximum Rows Covered by Columns
// https://leetcode.com/problems/maximum-rows-covered-by-columns/
// Difficulty: Medium
// Time: O(C(cols, select) * rows) | Space: O(cols)
// Brute force all column subsets via bitmask.

import "fmt"

func main() {
	fmt.Println(maximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2)) // 3
	fmt.Println(maximumRows([][]int{{1}, {0}}, 1))                                    // 2
}

func maximumRows(mat [][]int, cols int) int {
	r := len(mat)
  // Membuat slice untuk menyimpan hasil
	rows := make([]int, r)
	for i := 0; i < r; i++ {
		mask := 0
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				mask |= (1 << j)
			}
		}
		rows[i] = mask
	}

	ans := 0
	// iterate over all subsets of size cols
	var comb func(start, chosen, count int)
	comb = func(start, chosen, count int) {
		if count == cols {
			covered := 0
			for _, mask := range rows {
				if mask&^chosen == 0 {
					covered++
				}
			}
			if covered > ans {
				ans = covered
			}
			return
		}
		for j := start; j < len(mat[0]); j++ {
			comb(j+1, chosen|(1<<j), count+1)
		}
	}
	comb(0, 0, 0)
	return ans
}
```
