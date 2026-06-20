# 2135 — Count Words Obtained After Adding A Letter

## Deskripsi

**Soal:** [2135. Count Words Obtained After Adding A Letter](https://leetcode.com/problems/count-words-obtained-after-adding-a-letter/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * L)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Bitmask (representasi himpunan dengan bit)

**Fungsi Solusi:** `func wordCount(startWords []string, targetWords []string) int`

## Solusi Go

```go
package main

// LeetCode #2135: Count Words Obtained After Adding a Letter
// https://leetcode.com/problems/count-words-obtained-after-adding-a-letter/
// Difficulty: Medium
// Time: O(n * L) | Space: O(n)

import "fmt"

func wordCount(startWords []string, targetWords []string) int {
	// Convert start words to bitmasks
  // Membuat map untuk pencarian O(1): key → value
	startSet := make(map[int]bool)
	for _, w := range startWords {
		mask := 0
		for _, c := range w {
			mask |= 1 << (c - 'a')
		}
		startSet[mask] = true
	}

	count := 0
	for _, w := range targetWords {
		mask := 0
		for _, c := range w {
			mask |= 1 << (c - 'a')
		}
		// Try removing each character
		for _, c := range w {
			bit := 1 << (c - 'a')
			if startSet[mask^bit] {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", wordCount([]string{"ab", "a"}, []string{"abc", "abcd"}))
	// Expected: 1
}
```
