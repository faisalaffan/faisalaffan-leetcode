# 2881 — Create A New Column

## Deskripsi

**Soal:** [2881. Create A New Column](https://leetcode.com/problems/create-a-new-column/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2881: Create a New Column
// https://leetcode.com/problems/create-a-new-column/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we add a "grade" column computed from existing data.

import "fmt"

func main() {
	// LeetCode name: createBonusColumn
	fmt.Println(CreateANewColumn([][]int{{101, 15}, {102, 11}, {103, 20}}))
	// [[101 15 30] [102 11 22] [103 20 40]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: createBonusColumn
func CreateANewColumn(df [][]int) [][]int {
  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, len(df))
	for i, row := range df {
		// bonus = salary * 2
		result[i] = []int{row[0], row[1], row[1] * 2}
	}
	return result
}
```
