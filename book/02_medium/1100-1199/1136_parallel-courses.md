# 1136 — Parallel Courses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumSemesters(n int, relations [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS, Topological Sort

**Kompleksitas Waktu:** O(n + len(relations))  
**Kompleksitas Ruang:** O(n + len(relations))

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1136: Parallel Courses
// https://leetcode.com/problems/parallel-courses/
// Difficulty: Medium [Paid]

// topological sort (Kahn's algorithm) to find minimum semesters.

// Time: O(n + len(relations))
// Space: O(n + len(relations))

func minimumSemesters(n int, relations [][]int) int {
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n+1)
  // Alokasi slice integer
	indeg := make([]int, n+1)

	for _, r := range relations {
		adj[r[0]] = append(adj[r[0]], r[1])
		indeg[r[1]]++
	}

  // Alokasi slice integer
	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	semesters := 0
	taken := 0

	for len(queue) > 0 {
		semesters++
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			taken++
			for _, next := range adj[cur] {
				indeg[next]--
				if indeg[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}

	if taken != n {
		return -1
	}
	return semesters
}

func main() {
	fmt.Printf("%d (expected: 2)\n", minimumSemesters(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Printf("%d (expected: -1)\n", minimumSemesters(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	fmt.Printf("%d (expected: 1)\n", minimumSemesters(3, [][]int{}))
}
```
