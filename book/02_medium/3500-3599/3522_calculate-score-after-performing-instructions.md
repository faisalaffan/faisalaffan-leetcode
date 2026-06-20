# 3522 — Calculate Score After Performing Instructions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CalculateScoreAfterPerformingInstructions(ops []string, vals []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

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
  // Linear scan O(n)
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
