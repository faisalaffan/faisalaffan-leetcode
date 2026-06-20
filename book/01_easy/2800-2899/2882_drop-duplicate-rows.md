# 2882 — Drop Duplicate Rows

## Deskripsi

**Soal:** [2882. Drop Duplicate Rows](https://leetcode.com/problems/drop-duplicate-rows/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2882: Drop Duplicate Rows
// https://leetcode.com/problems/drop-duplicate-rows/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we remove rows with duplicate emails.

import "fmt"

func main() {
	// LeetCode name: dropDuplicateEmails
	fmt.Println(DropDuplicateRows([][]string{{"1", "a@b.com"}, {"2", "c@d.com"}, {"3", "a@b.com"}}))
	// [[1 a@b.com] [2 c@d.com]]

	fmt.Println(DropDuplicateRows([][]string{{"1", "x@y.com"}, {"2", "x@y.com"}, {"3", "x@y.com"}}))
	// [[1 x@y.com]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: dropDuplicateEmails
func DropDuplicateRows(df [][]string) [][]string {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[string]bool)
	result := [][]string{}
	for _, row := range df {
		email := row[1]
		if !seen[email] {
			seen[email] = true
			result = append(result, row)
		}
	}
	return result
}
```
