# 3817 — Good Indices In A Digit String

## Deskripsi

**Soal:** [3817. Good Indices In A Digit String](https://leetcode.com/problems/good-indices-in-a-digit-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N * L) where L <= 6  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func GoodIndicesInADigitString(s string) []int`

> **Ide Kunci:** For each index, check if a substring ending at i equals decimal representation of i.

## Solusi Go

```go
package main

// LeetCode #3817: Good Indices in a Digit String
// https://leetcode.com/problems/good-indices-in-a-digit-string/
// Difficulty: Medium [Paid]
// Time: O(N * L) where L <= 6 | Space: O(1)
// Approach: For each index, check if a substring ending at i equals decimal representation of i.

import (
	"fmt"
	"strconv"
)

func GoodIndicesInADigitString(s string) []int {
	n := len(s)
	ans := []int{}

	for i := 0; i < n; i++ {
		rep := strconv.Itoa(i)
		l := len(rep)
		if i-l+1 >= 0 && s[i-l+1:i+1] == rep {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(GoodIndicesInADigitString("0234567890112")) // Expected: [0 11 12]

	// Example 2
	fmt.Println(GoodIndicesInADigitString("01234")) // Expected: [0 1 2 3 4]

	// Example 3
	fmt.Println(GoodIndicesInADigitString("12345")) // Expected: []
}
```
