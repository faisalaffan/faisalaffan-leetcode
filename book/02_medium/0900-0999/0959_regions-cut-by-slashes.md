# 0959 — Regions Cut By Slashes

## Deskripsi

**Soal:** [0959. Regions Cut By Slashes](https://leetcode.com/problems/regions-cut-by-slashes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2 * α(n^2))  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** —

**Fungsi Solusi:** `func regionsBySlashes(grid []string) int`

## Solusi Go

```go
package main

// LeetCode #959: Regions Cut By Slashes
// https://leetcode.com/problems/regions-cut-by-slashes/
// Difficulty: Medium

import "fmt"

// Time: O(n^2 * α(n^2)) | Space: O(n^2)
func regionsBySlashes(grid []string) int {
	n := len(grid)
	size := n * n * 4
  // Membuat slice untuk menyimpan hasil
	parent := make([]int, size)
  // Iterasi seluruh elemen
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[rb] = ra
			size--
		}
	}

	for i, row := range grid {
		for j, ch := range row {
			base := (i*n + j) * 4
			// Connect internal triangles
			if ch == '/' {
				union(base, base+3)
				union(base+1, base+2)
			} else if ch == '\\' {
				union(base, base+1)
				union(base+2, base+3)
			} else {
				union(base, base+1)
				union(base+1, base+2)
				union(base+2, base+3)
			}
			// Connect with right neighbor
			if j+1 < n {
				union(base+1, (base+4)+3)
			}
			// Connect with bottom neighbor
			if i+1 < n {
				union(base+2, (base+4*n))
			}
		}
	}

	return size
}

func main() {
	fmt.Println(regionsBySlashes([]string{" /", "/ "}))
	fmt.Println(regionsBySlashes([]string{" /", "  "}))
	fmt.Println(regionsBySlashes([]string{"/\\", "\\/"}))
}
```
