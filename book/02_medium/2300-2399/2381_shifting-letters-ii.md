# 2381 — Shifting Letters Ii

## Deskripsi

**Soal:** [2381. Shifting Letters Ii](https://leetcode.com/problems/shifting-letters-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2381: Shifting Letters II
// https://leetcode.com/problems/shifting-letters-ii/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)
// Difference array to apply range shifts efficiently.

import "fmt"

func main() {
	fmt.Println(shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}})) // "ace"
	fmt.Println(shiftingLetters("dztz", [][]int{{0, 0, 0}, {1, 1, 1}}))       // "catz"
}

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
  // Membuat slice untuk menyimpan hasil
	diff := make([]int, n+1)
	for _, sh := range shifts {
		start, end, dir := sh[0], sh[1], sh[2]
		if dir == 1 {
			diff[start]++
			diff[end+1]--
		} else {
			diff[start]--
			diff[end+1]++
		}
	}

	cur := 0
  // Membuat slice untuk menyimpan hasil
	res := make([]byte, n)
	for i, ch := range s {
		cur += diff[i]
		shift := ((int(ch-'a')+cur)%26 + 26) % 26
		res[i] = byte('a' + shift)
	}
	return string(res)
}
```
