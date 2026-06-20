# 1861 — Rotating The Box

## Deskripsi

**Soal:** [1861. Rotating The Box](https://leetcode.com/problems/rotating-the-box/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n), Space: O(m*n) for result  
**Kompleksitas Ruang:** O(m*n) for result

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1861: Rotating the Box
// https://leetcode.com/problems/rotating-the-box/
// Difficulty: Medium

import "fmt"

func main() {
	box1 := [][]byte{{'#', '.', '#'}}
	fmt.Println(RotateTheBox(box1))

	box2 := [][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}
	fmt.Println(RotateTheBox(box2))

	box3 := [][]byte{{'#', '#', '*', '.', '*', '.'},
		{'#', '#', '#', '*', '.', '.'},
		{'#', '#', '#', '.', '#', '.'}}
	fmt.Println(RotateTheBox(box3))
}

// Time: O(m*n), Space: O(m*n) for result
func RotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])

	// Apply gravity to each row (stones fall to the right)
	for i := 0; i < m; i++ {
		empty := n - 1
		for j := n - 1; j >= 0; j-- {
			if box[i][j] == '*' {
				empty = j - 1
			} else if box[i][j] == '#' {
				box[i][j] = '.'
				box[i][empty] = '#'
				empty--
			}
		}
	}

	// Rotate 90 degrees clockwise
  // Membuat slice 2D untuk DP/tabel
	result := make([][]byte, n)
  // Iterasi seluruh elemen
	for i := range result {
		result[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[j][m-1-i] = box[i][j]
		}
	}
	return result
}
```
