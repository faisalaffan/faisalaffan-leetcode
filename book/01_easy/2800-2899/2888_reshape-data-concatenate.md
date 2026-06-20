# 2888 — Reshape Data Concatenate

## Deskripsi

**Soal:** [2888. Reshape Data Concatenate](https://leetcode.com/problems/reshape-data-concatenate/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n+m)  
**Kompleksitas Ruang:** O(n+m)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2888: Reshape Data: Concatenate
// https://leetcode.com/problems/reshape-data-concatenate/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we concatenate two dataframes vertically.

import "fmt"

func main() {
	// LeetCode name: concatenateDataFrames
	fmt.Println(ReshapeDataConcatenate([][]int{{1, 15}, {2, 11}}, [][]int{{3, 12}, {4, 14}}))
	// [[1 15] [2 11] [3 12] [4 14]]
}

// Time: O(n+m) | Space: O(n+m)
// LeetCode submission name: concatenateDataFrames
func ReshapeDataConcatenate(df1, df2 [][]int) [][]int {
  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0, len(df1)+len(df2))
	result = append(result, df1...)
	result = append(result, df2...)
	return result
}
```
