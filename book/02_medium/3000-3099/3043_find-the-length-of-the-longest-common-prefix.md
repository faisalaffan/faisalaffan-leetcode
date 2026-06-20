# 3043 — Find The Length Of The Longest Common Prefix

## Deskripsi

**Soal:** [3043. Find The Length Of The Longest Common Prefix](https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n*logM + m*logM)  
**Kompleksitas Ruang:** O(n*logM)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3043: Find the Length of the Longest Common Prefix
// https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/
// Difficulty: Medium
// Time: O(n*logM + m*logM) | Space: O(n*logM)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
	fmt.Println(longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4}))
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixes := map[int]bool{}
	for _, x := range arr1 {
		for x > 0 {
			prefixes[x] = true
			x /= 10
		}
	}
	ans := 0
	for _, x := range arr2 {
		for x > 0 {
			if prefixes[x] {
				if len(strconv.Itoa(x)) > ans {
					ans = len(strconv.Itoa(x))
				}
				break
			}
			x /= 10
		}
	}
	return ans
}
```
