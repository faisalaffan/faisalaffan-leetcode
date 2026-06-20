# 2133 — Check If Every Row And Column Contains All Numbers

## Deskripsi

**Soal:** [2133. Check If Every Row And Column Contains All Numbers](https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2133: Check if Every Row and Column Contains All Numbers
// https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfEveryRowAndColumnContainsAllNumbers([][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}})) // true
	fmt.Println(CheckIfEveryRowAndColumnContainsAllNumbers([][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}})) // false
}

// Time: O(n^2), Space: O(n)
func CheckIfEveryRowAndColumnContainsAllNumbers(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
  // Membuat slice untuk menyimpan hasil
		rowSet := make([]bool, n+1)
  // Membuat slice untuk menyimpan hasil
		colSet := make([]bool, n+1)
		for j := 0; j < n; j++ {
			if matrix[i][j] < 1 || matrix[i][j] > n || rowSet[matrix[i][j]] {
				return false
			}
			rowSet[matrix[i][j]] = true

			if matrix[j][i] < 1 || matrix[j][i] > n || colSet[matrix[j][i]] {
				return false
			}
			colSet[matrix[j][i]] = true
		}
	}
	return true
}
```
