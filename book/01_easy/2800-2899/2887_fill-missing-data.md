# 2887 — Fill Missing Data

## Deskripsi

**Soal:** [2887. Fill Missing Data](https://leetcode.com/problems/fill-missing-data/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2887: Fill Missing Data
// https://leetcode.com/problems/fill-missing-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we fill missing (zero/empty) quantity values with 0.

import "fmt"

func main() {
	// LeetCode name: fillMissingValues
	// Input: [name, quantity]. Empty string means missing.
	fmt.Println(FillMissingData([][]string{{"A", "10"}, {"B", ""}, {"C", "5"}, {"D", ""}}))
	// [[A 10] [B 0] [C 5] [D 0]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: fillMissingValues
func FillMissingData(df [][]string) [][]string {
  // Membuat slice 2D untuk DP/tabel
	result := make([][]string, len(df))
	for i, row := range df {
		if row[1] == "" {
			result[i] = []string{row[0], "0"}
		} else {
			result[i] = row
		}
	}
	return result
}
```
