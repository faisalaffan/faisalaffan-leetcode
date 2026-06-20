# 0844 — Backspace String Compare

## Deskripsi

**Soal:** [0844. Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #844: Backspace String Compare
// https://leetcode.com/problems/backspace-string-compare/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
	fmt.Println(backspaceCompare("ab##", "c#d#")) // true
	fmt.Println(backspaceCompare("a#c", "b"))     // false
	fmt.Println(backspaceCompare("a##c", "#a#c")) // true
}

// backspaceCompare checks if two strings are equal after applying backspace.
// Time: O(n + m). Space: O(1).
func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	skipS, skipT := 0, 0
	for i >= 0 || j >= 0 {
		// Find next valid char in s
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		// Find next valid char in t
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
```
