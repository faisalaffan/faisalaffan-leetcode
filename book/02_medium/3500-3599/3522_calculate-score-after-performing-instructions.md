# 3522 — Calculate Score After Performing Instructions

## Deskripsi

**Soal:** [3522. Calculate Score After Performing Instructions](https://leetcode.com/problems/calculate-score-after-performing-instructions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3522: Calculate Score After Performing Instructions
// https://leetcode.com/problems/calculate-score-after-performing-instructions/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", CalculateScoreAfterPerformingInstructions([]string{"add", "add", "sub"}, []int{5, 3, 2}))
	// Test case 2
	fmt.Println("Test 2:", CalculateScoreAfterPerformingInstructions([]string{"add", "mul", "add"}, []int{1, 2, 3}))
	// Test case 3
	fmt.Println("Test 3:", CalculateScoreAfterPerformingInstructions([]string{"add"}, []int{10}))
}

func CalculateScoreAfterPerformingInstructions(ops []string, vals []int) int {
	score := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(ops) && i < len(vals); i++ {
		switch ops[i] {
		case "add":
			score += vals[i]
		case "sub":
			score -= vals[i]
		case "mul":
			score *= vals[i]
		case "div":
			if vals[i] != 0 {
				score /= vals[i]
			}
		}
	}
	return score
}
```
