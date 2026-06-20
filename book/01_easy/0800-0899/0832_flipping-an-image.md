# 0832 — Flipping An Image

## Deskripsi

**Soal:** [0832. Flipping An Image](https://leetcode.com/problems/flipping-an-image/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(m*n). Space: O(1) in-place.  
**Kompleksitas Ruang:** O(1) in-place.

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #832: Flipping an Image
// https://leetcode.com/problems/flipping-an-image/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}})) // [[1,0,0],[0,1,0],[1,1,1]]
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}

// flipAndInvertImage flips the image horizontally then inverts it.
// Time: O(m*n). Space: O(1) in-place.
func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		l, r := 0, len(row)-1
		for l <= r {
			row[l], row[r] = 1-row[r], 1-row[l]
			l++
			r--
		}
	}
	return image
}
```
