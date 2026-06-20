# 0014 — Longest Common Prefix

## Deskripsi

**Soal:** [0014. Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n*m)  
**Kompleksitas Ruang:** O(1) where n=len(strs), m=len(shortest string)

**Algoritma:** —

**Fungsi Solusi:** `func LongestCommonPrefix(strs []string) string`

## Solusi Go

```go
package main

// LeetCode #14: Longest Common Prefix
// https://leetcode.com/problems/longest-common-prefix/
// Difficulty: Easy

import "fmt"

// Time: O(n*m) | Space: O(1) where n=len(strs), m=len(shortest string)
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(LongestCommonPrefix([]string{"a"}))
}
```
