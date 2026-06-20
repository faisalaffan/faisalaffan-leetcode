# 1284 — Minimum Number Of Flips To Convert Binary Matrix To Zero Matrix

## Deskripsi

**Soal:** [1284. Minimum Number Of Flips To Convert Binary Matrix To Zero Matrix](https://leetcode.com/problems/minimum-number-of-flips-to-convert-binary-matrix-to-zero-matrix/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** BFS (Breadth-First Search / pencarian lebar), Bitmask (representasi himpunan dengan bit)

**Fungsi Solusi:** `func minFlips(mat [][]int) int`

> **Ide Kunci:** BFS over bitmask states.

## Solusi Go

```go
package main

// LeetCode #1284: Minimum Number of Flips to Convert Binary Matrix to Zero Matrix
// https://leetcode.com/problems/minimum-number-of-flips-to-convert-binary-matrix-to-zero-matrix/
// Difficulty: Hard
//
// Approach: BFS over bitmask states.
// Max matrix dimension is 3x3 => at most 9 bits. Each state is an integer
// bitmask representing the matrix. BFS from the initial state to 0,
// flipping each cell (and its 4-direction neighbors) per step.

import "fmt"

func minFlips(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	start := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				start |= 1 << (i*n + j)
			}
		}
	}
	if start == 0 {
		return 0
	}

	dirs := [][]int{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}
  // Membuat map untuk pencarian O(1): key → value
	visited := make(map[int]bool)
	q := []int{start}
	visited[start] = true
	steps := 0

	for len(q) > 0 {
		steps++
		for sz := len(q); sz > 0; sz-- {
			cur := q[0]
			q = q[1:]
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					nxt := cur
					for _, d := range dirs {
						ni, nj := i+d[0], j+d[1]
						if ni >= 0 && ni < m && nj >= 0 && nj < n {
							nxt ^= 1 << (ni*n + nj)
						}
					}
					if nxt == 0 {
						return steps
					}
					if !visited[nxt] {
						visited[nxt] = true
						q = append(q, nxt)
					}
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minFlips([][]int{{0, 0}, {0, 1}}))                      // 3
	fmt.Println(minFlips([][]int{{0}}))                                  // 0
	fmt.Println(minFlips([][]int{{1, 0, 0}, {1, 0, 0}}))                // -1
	fmt.Println(minFlips([][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 0}}))    // -1
	fmt.Println(minFlips([][]int{{1}}))                                  // 1
}
```
