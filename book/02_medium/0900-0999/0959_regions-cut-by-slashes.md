# 0959 — Regions Cut By Slashes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func regionsBySlashes(grid []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Union-Find

**Waktu:** O(n^2 * α(n^2))  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **Union-Find** — sering muncul di interview!

## 💻 Solusi Go

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
  // Alokasi slice
	parent := make([]int, size)
  // Range loop
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
