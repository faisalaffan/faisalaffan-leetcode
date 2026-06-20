# 0959 — Regions Cut By Slashes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func regionsBySlashes(grid []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** O(n^2 * α(n^2))  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Alokasi slice integer
	parent := make([]int, size)
  // Range loop: iterasi dengan indeks + nilai
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
