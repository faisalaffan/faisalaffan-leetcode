# 1758 — Minimum Changes To Make Alternating Binary String

## Deskripsi

**Soal:** [1758. Minimum Changes To Make Alternating Binary String](https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MinOperations(s string) int`

## Solusi Go

```go
package main

// LeetCode #1758: Minimum Changes to Make Alternating Binary String
// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MinOperations(s string) int {
	changes := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		expected := byte('0' + i%2)
		if s[i] != expected {
			changes++
		}
	}
	if changes < len(s)-changes {
		return changes
	}
	return len(s) - changes
}

func main() {
	fmt.Println(MinOperations("0100"))
	fmt.Println(MinOperations("10"))
	fmt.Println(MinOperations("1111"))
}
```
