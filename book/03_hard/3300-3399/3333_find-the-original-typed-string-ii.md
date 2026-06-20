# 3333 — Find The Original Typed String Ii

## Deskripsi

**Soal:** [3333. Find The Original Typed String Ii](https://leetcode.com/problems/find-the-original-typed-string-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Group consecutive same characters. For each group of

## Solusi Go

```go
package main

// LeetCode #3333: Find the Original Typed String II
// https://leetcode.com/problems/find-the-original-typed-string-ii/
// Difficulty: Hard
//
// Alice typed a string word but some characters may be long-pressed
// (the character is repeated). Given the final string and k, count
// possible original strings where no character was typed more than
// k times consecutively.
//
// Approach: Group consecutive same characters. For each group of
// length len, the original could have any length from 1 to min(len,k).
// Multiply possibilities across groups.

import "fmt"

func main() {
	// Example 1
	fmt.Println(possibleStringCount("aabbccdd", 2))
	// Example 2
	fmt.Println(possibleStringCount("aaaa", 2))
	// Edge: single char
	fmt.Println(possibleStringCount("a", 5))
}

const STR_MOD = 1000000007

func possibleStringCount(word string, k int) int {
	n := len(word)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// Count runs
	var runs []int
	i := 0
	for i < n {
		j := i
		for j < n && word[j] == word[i] {
			j++
		}
		runs = append(runs, j-i)
		i = j
	}

	ans := 1
	for _, r := range runs {
		// Original could have length 1 to min(r, k)
		options := r
		if options > k {
			options = k
		}
		ans = (ans * options) % STR_MOD
	}

	return ans
}
```
