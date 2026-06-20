# 0925 — Long Pressed Name

## Deskripsi

**Soal:** [0925. Long Pressed Name](https://leetcode.com/problems/long-pressed-name/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #925: Long Pressed Name
// https://leetcode.com/problems/long-pressed-name/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))    // true
	fmt.Println(isLongPressedName("saeed", "ssaaedd"))  // false
	fmt.Println(isLongPressedName("leelee", "lleeelee")) // true
	fmt.Println(isLongPressedName("alex", "aaleexa"))   // false
}

// isLongPressedName checks if typed is a long-pressed version of name.
// Time: O(n + m). Space: O(1).
func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		// Count occurrences in name
		c1 := 0
		ch := name[i]
		for i < len(name) && name[i] == ch {
			i++
			c1++
		}
		// Count occurrences in typed
		c2 := 0
		for j < len(typed) && typed[j] == ch {
			j++
			c2++
		}
		if c2 < c1 {
			return false
		}
	}
	return i == len(name) && j == len(typed)
}
```
