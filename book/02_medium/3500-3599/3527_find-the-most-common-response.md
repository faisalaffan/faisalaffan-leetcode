# 3527 — Find The Most Common Response

## Deskripsi

**Soal:** [3527. Find The Most Common Response](https://leetcode.com/problems/find-the-most-common-response/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3527: Find the Most Common Response
// https://leetcode.com/problems/find-the-most-common-response/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindTheMostCommonResponse([]string{"A", "B", "A", "C", "B", "A"}))
	// Test case 2
	fmt.Println("Test 2:", FindTheMostCommonResponse([]string{"X", "Y", "Z"}))
	// Test case 3
	fmt.Println("Test 3:", FindTheMostCommonResponse([]string{"M", "M", "M"}))
}

func FindTheMostCommonResponse(responses []string) string {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[string]int)
	maxFreq := 0
	mostCommon := ""
	for _, r := range responses {
		freq[r]++
		if freq[r] > maxFreq || (freq[r] == maxFreq && (mostCommon == "" || r < mostCommon)) {
			maxFreq = freq[r]
			mostCommon = r
		}
	}
	return mostCommon
}
```
