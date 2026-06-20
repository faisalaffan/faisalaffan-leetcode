# 1111 — Maximum Nesting Depth Of Two Valid Parentheses Strings

## Deskripsi

**Soal:** [1111. Maximum Nesting Depth Of Two Valid Parentheses Strings](https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

> **Ide Kunci:** Assign '(' to group A or B based on even/odd depth.

## Solusi Go

```go
package main

// LeetCode #1111: Maximum Nesting Depth of Two Valid Parentheses Strings
// https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/
// Difficulty: Medium
//
// Approach: Assign '(' to group A or B based on even/odd depth.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxDepthAfterSplit("(()())")) // [0,1,1,1,1,0] or similar
	fmt.Println(maxDepthAfterSplit("()(())()")) // [0,0,0,1,1,0,0,0]
}

func maxDepthAfterSplit(seq string) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(seq))
	depth := 0

	for i, c := range seq {
		if c == '(' {
			depth++
			result[i] = depth % 2
		} else {
			result[i] = depth % 2
			depth--
		}
	}

	return result
}
```
