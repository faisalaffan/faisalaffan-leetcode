# 1409 — Queries On A Permutation With Key

## Deskripsi

**Soal:** [1409. Queries On A Permutation With Key](https://leetcode.com/problems/queries-on-a-permutation-with-key/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n) where m = len(queries), n = m (since P has m elements)  
**Kompleksitas Ruang:** O(n) for the permutation

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1409: Queries on a Permutation With Key
// https://leetcode.com/problems/queries-on-a-permutation-with-key/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5)) // [2,1,2,1]

	// Test case 2
	fmt.Println(processQueries([]int{4, 1, 2, 2}, 4)) // [3,1,2,0]

	// Test case 3
	fmt.Println(processQueries([]int{7, 5, 5, 8, 3}, 8)) // [6,5,0,7,5]
}

// Time: O(m*n) where m = len(queries), n = m (since P has m elements)
// Space: O(n) for the permutation
func processQueries(queries []int, m int) []int {
	// Build permutation P = [1, 2, ..., m]
  // Membuat slice untuk menyimpan hasil
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = i + 1
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(queries))

	for idx, q := range queries {
		// Find position of q in P
		pos := 0
		for p[pos] != q {
			pos++
		}
		result[idx] = pos

		// Move q to front by shifting elements before it
		for i := pos; i > 0; i-- {
			p[i] = p[i-1]
		}
		p[0] = q
	}

	return result
}
```
