# 1292 — Maximum Side Length Of A Square With Sum Less Than Or Equal To Threshold

## Deskripsi

**Soal:** [1292. Maximum Side Length Of A Square With Sum Less Than Or Equal To Threshold](https://leetcode.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n * log(min(m,n)))  
**Kompleksitas Ruang:** O(m*n)

**Algoritma:** Binary Search (pencarian biner), Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func maxSideLength(mat [][]int, threshold int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1292: Maximum Side Length of a Square with Sum Less than or Equal to Threshold
// https://leetcode.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/
// Difficulty: Medium

// Prefix sum matrix + binary search on side length.
// For each square, sum = prefix[r+s][c+s] - prefix[r][c+s] - prefix[r+s][c] + prefix[r][c].

// Time: O(m * n * log(min(m,n)))
// Space: O(m*n)

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
  // Membuat slice 2D untuk DP/tabel
	prefix := make([][]int, m+1)
  // Iterasi seluruh elemen
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			prefix[i][j] = mat[i-1][j-1] + prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1]
		}
	}

	maxSide := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for s := maxSide + 1; s <= m-i+1 && s <= n-j+1; s++ {
				sum := prefix[i+s-1][j+s-1] - prefix[i-1][j+s-1] - prefix[i+s-1][j-1] + prefix[i-1][j-1]
				if sum <= threshold {
					if s > maxSide {
						maxSide = s
					}
				} else {
					break
				}
			}
		}
	}

	return maxSide
}

func main() {
	fmt.Printf("%d (expected: 2)\n",
		maxSideLength([][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, 4))

	fmt.Printf("%d (expected: 3)\n",
		maxSideLength([][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, 12))
}
```
