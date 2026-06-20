# 0806 — Number Of Lines To Write String

## Deskripsi

**Soal:** [0806. Number Of Lines To Write String](https://leetcode.com/problems/number-of-lines-to-write-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #806: Number of Lines To Write String
// https://leetcode.com/problems/number-of-lines-to-write-string/
// Difficulty: Easy

import "fmt"

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	fmt.Println(numberOfLines(widths, "abcdefghijklmnopqrstuvwxyz")) // [3, 60]

	widths2 := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	fmt.Println(numberOfLines(widths2, "bbbcccdddaaa")) // [2, 4]
}

// numberOfLines returns the lines and last line width needed to write the string.
// Time: O(n). Space: O(1).
func numberOfLines(widths []int, s string) []int {
	lines, currentWidth := 1, 0
	for _, c := range s {
		w := widths[c-'a']
		if currentWidth+w > 100 {
			lines++
			currentWidth = w
		} else {
			currentWidth += w
		}
	}
	return []int{lines, currentWidth}
}
```
