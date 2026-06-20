# 2947 — Count Beautiful Substrings I

## Deskripsi

**Soal:** [2947. Count Beautiful Substrings I](https://leetcode.com/problems/count-beautiful-substrings-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2947: Count Beautiful Substrings I
// https://leetcode.com/problems/count-beautiful-substrings-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(beautifulSubstrings("baeyh", 2))
	fmt.Println(beautifulSubstrings("ab", 1))
	fmt.Println(beautifulSubstrings("a", 1))
}

func beautifulSubstrings(s string, k int) (ans int) {
	n := len(s)
	vowels := [26]bool{}
	for _, c := range "aeiou" {
		vowels[c-'a'] = true
	}
	for i := 0; i < n; i++ {
		v := 0
		for j := i; j < n; j++ {
			if vowels[s[j]-'a'] {
				v++
			}
			c := j - i + 1 - v
			if v == c && v*c%k == 0 {
				ans++
			}
		}
	}
	return
}
```
