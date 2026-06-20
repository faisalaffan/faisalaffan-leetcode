# 0520 — Detect Capital

## Deskripsi

**Soal:** [0520. Detect Capital](https://leetcode.com/problems/detect-capital/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func DetectCapital(word string) bool`

## Solusi Go

```go
package main

// LeetCode #520: Detect Capital
// https://leetcode.com/problems/detect-capital/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func DetectCapital(word string) bool {
	upperCount := 0
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			upperCount++
		}
	}
	return upperCount == len(word) || upperCount == 0 || (upperCount == 1 && word[0] >= 'A' && word[0] <= 'Z')
}

func main() {
	fmt.Println(DetectCapital("USA"))
	fmt.Println(DetectCapital("FlaG"))
	fmt.Println(DetectCapital("Google"))
}
```
