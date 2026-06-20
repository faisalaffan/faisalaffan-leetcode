# 0684 — Redundant Connection

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findRedundantConnection(edges [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** O(n * alpha(n))  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #684: Redundant Connection
// https://leetcode.com/problems/redundant-connection/
// Difficulty: Medium
// Time: O(n * alpha(n))
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
  // Alokasi slice integer
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		px, py := find(x), find(y)
		if px == py {
			return false
		}
		parent[px] = py
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
```
