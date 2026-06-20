# 0947 — Most Stones Removed With Same Row Or Column

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeStones(stones [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Union-Find (DSU)

**Kompleksitas Waktu:** O(n * α(n))  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #947: Most Stones Removed with Same Row or Column
// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
// Difficulty: Medium

import "fmt"

// Time: O(n * α(n)) | Space: O(n)
func removeStones(stones [][]int) int {
	n := len(stones)
  // Alokasi slice integer
	parent := make([]int, n)
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
		}
	}

	// Union stones sharing row or column
  // Membuat map (HashMap) — pencarian O(1)
	rowMap := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	colMap := make(map[int]int)
	for i, s := range stones {
		if v, ok := rowMap[s[0]]; ok {
			union(i, v)
		} else {
			rowMap[s[0]] = i
		}
		if v, ok := colMap[s[1]]; ok {
			union(i, v)
		} else {
			colMap[s[1]] = i
		}
	}

	// Count connected components
	components := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			components++
		}
	}

	return n - components
}

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 0}}))
}
```
