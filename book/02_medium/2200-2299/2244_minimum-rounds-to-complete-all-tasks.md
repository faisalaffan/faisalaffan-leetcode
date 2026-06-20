# 2244 — Minimum Rounds To Complete All Tasks

## Deskripsi

**Soal:** [2244. Minimum Rounds To Complete All Tasks](https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func minimumRounds(tasks []int) int`

## Solusi Go

```go
package main

// LeetCode #2244: Minimum Rounds to Complete All Tasks
// https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumRounds(tasks []int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, t := range tasks {
		freq[t]++
	}

	rounds := 0
	for _, c := range freq {
		if c == 1 {
			return -1
		}
		// 3*c + 2*(c/3 remainder)
		rounds += c / 3
		if c%3 != 0 {
			rounds++
		}
	}
	return rounds
}

func main() {
	// Test case 1
	fmt.Println(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
	// Expected: 4

	// Test case 2
	fmt.Println(minimumRounds([]int{2, 3, 3}))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumRounds([]int{5, 5, 5, 5}))
	// Expected: 2
}
```
