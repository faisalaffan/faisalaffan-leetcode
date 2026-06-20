# 2782 — Number Of Unique Categories

## Deskripsi

**Soal:** [2782. Number Of Unique Categories](https://leetcode.com/problems/number-of-unique-categories/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func NumberOfUniqueCategories(categories []string) int`

## Solusi Go

```go
package main

// LeetCode #2782: Number of Unique Categories
// https://leetcode.com/problems/number-of-unique-categories/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func NumberOfUniqueCategories(categories []string) int {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[string]bool)
	for _, c := range categories {
		seen[c] = true
	}
	return len(seen)
}

func main() {
	fmt.Println(NumberOfUniqueCategories([]string{"a", "b", "a", "c"}))
	fmt.Println(NumberOfUniqueCategories([]string{"x", "x", "x"}))
	fmt.Println(NumberOfUniqueCategories([]string{}))
}
```
