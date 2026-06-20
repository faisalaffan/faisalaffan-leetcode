# 1207 — Unique Number Of Occurrences

## Deskripsi

**Soal:** [1207. Unique Number Of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1207: Unique Number of Occurrences
// https://leetcode.com/problems/unique-number-of-occurrences/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})) // true
	fmt.Println(uniqueOccurrences([]int{1, 2}))             // false
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})) // true
}

// LeetCode submission: uniqueOccurrences
func uniqueOccurrences(arr []int) bool {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)
	for _, f := range freq {
		if seen[f] {
			return false
		}
		seen[f] = true
	}
	return true
}
```
