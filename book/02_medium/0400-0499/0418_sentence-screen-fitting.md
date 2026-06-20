# 0418 — Sentence Screen Fitting

## Deskripsi

**Soal:** [0418. Sentence Screen Fitting](https://leetcode.com/problems/sentence-screen-fitting/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(rows * avgWordLen)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func wordsTyping(sentence []string, rows int, cols int) int`

## Solusi Go

```go
package main

// LeetCode #418: Sentence Screen Fitting
// https://leetcode.com/problems/sentence-screen-fitting/
// Difficulty: Medium [Paid]
// Time: O(rows * avgWordLen) | Space: O(1)

import "fmt"

func wordsTyping(sentence []string, rows int, cols int) int {
	s := ""
	for _, w := range sentence {
		s += w + " "
	}

	n := len(s)
	start := 0

	for i := 0; i < rows; i++ {
		start += cols

		// If next char is a space, we can fit perfectly
		if s[start%n] == ' ' {
			start++
		} else {
			// Backtrack to nearest space
			for start > 0 && s[(start-1)%n] != ' ' {
				start--
			}
		}
	}
	return start / n
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wordsTyping([]string{"hello", "world"}, 2, 8))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", wordsTyping([]string{"a", "bcd", "e"}, 3, 6))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", wordsTyping([]string{"i", "had", "apple", "pie"}, 4, 5))
	// Expected: 1
}
```
