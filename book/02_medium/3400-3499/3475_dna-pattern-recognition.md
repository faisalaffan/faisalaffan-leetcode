# 3475 — Dna Pattern Recognition

## Deskripsi

**Soal:** [3475. Dna Pattern Recognition](https://leetcode.com/problems/dna-pattern-recognition/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3475: DNA Pattern Recognition
// https://leetcode.com/problems/dna-pattern-recognition/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", DnaPatternRecognition("ACGTACGT", "ACGT"))
	// Test case 2
	fmt.Println("Test 2:", DnaPatternRecognition("AAAA", "AA"))
	// Test case 3
	fmt.Println("Test 3:", DnaPatternRecognition("ACGT", "TGCA"))
}

func DnaPatternRecognition(dna string, pattern string) int {
	// Count occurrences of pattern in DNA string
	if len(pattern) == 0 {
		return 0
	}
	count := 0
	for i := 0; i <= len(dna)-len(pattern); i++ {
		match := true
		for j := 0; j < len(pattern); j++ {
			if dna[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	return count
}
```
