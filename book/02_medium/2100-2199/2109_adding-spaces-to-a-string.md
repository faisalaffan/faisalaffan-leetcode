# 2109 — Adding Spaces To A String

## Deskripsi

**Soal:** [2109. Adding Spaces To A String](https://leetcode.com/problems/adding-spaces-to-a-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m)

**Algoritma:** —

**Fungsi Solusi:** `func addSpaces(s string, spaces []int) string`

## Solusi Go

```go
package main

// LeetCode #2109: Adding Spaces to a String
// https://leetcode.com/problems/adding-spaces-to-a-string/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n + m)

import (
	"fmt"
	"strings"
)

func addSpaces(s string, spaces []int) string {
	var result strings.Builder
	spaceIdx := 0
	n := len(s)

	for i := 0; i < n; i++ {
		if spaceIdx < len(spaces) && i == spaces[spaceIdx] {
			result.WriteByte(' ')
			spaceIdx++
		}
		result.WriteByte(s[i])
	}

	return result.String()
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
	// Expected: "Leetcode Helps Me Learn"

	// Test case 2
	fmt.Println("Test 2:", addSpaces("icodeinpython", []int{1, 5, 7, 9}))
	// Expected: "i code in py thon"

	// Test case 3
	fmt.Println("Test 3:", addSpaces("spacing", []int{0}))
	// Expected: " spacing"
}
```
