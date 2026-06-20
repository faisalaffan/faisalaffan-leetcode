# 2643 — Row With Maximum Ones

## Deskripsi

**Soal:** [2643. Row With Maximum Ones](https://leetcode.com/problems/row-with-maximum-ones/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2643: Row With Maximum Ones
// https://leetcode.com/problems/row-with-maximum-ones/
// Difficulty: Easy
// Time: O(m * n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(RowWithMaximumOnes([][]int{{0, 1}, {1, 0}}))
	fmt.Println(RowWithMaximumOnes([][]int{{0, 0, 0}, {0, 1, 1}}))
}

func RowWithMaximumOnes(mat [][]int) []int {
	maxRow, maxCount := 0, 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(mat); i++ {
		count := 0
		for _, val := range mat[i] {
			if val == 1 {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			maxRow = i
		}
	}
	return []int{maxRow, maxCount}
}
```
