# 3798 — Largest Even Number

## Deskripsi

**Soal:** [3798. Largest Even Number](https://leetcode.com/problems/largest-even-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3798: Largest Even Number
// https://leetcode.com/problems/largest-even-number/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LargestEvenNumber("1112"))
	fmt.Println(LargestEvenNumber("221"))
	fmt.Println(LargestEvenNumber("1"))
}

// Time: O(n)
// Space: O(n)
func LargestEvenNumber(s string) string {
	i := len(s)
	for i > 0 && s[i-1] == '1' {
		i--
	}
	return s[:i]
}
```
