# 2158 — Amount Of New Area Painted Each Day

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func amountPainted(paint [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2158: Amount of New Area Painted Each Day
// https://leetcode.com/problems/amount-of-new-area-painted-each-day/
// Difficulty: Hard [Paid]
//
// Given intervals [start, end) painted each day, compute the new area painted
// each day (not previously painted). Use a DSU/union-find structure to skip
// already-painted cells, achieving near O(N) amortized time.

import "fmt"

func main() {
	// Example 1
	fmt.Println(amountPainted([][]int{{1, 4}, {4, 7}, {5, 8}}))

	// Example 2
	fmt.Println(amountPainted([][]int{{1, 5}, {2, 4}}))

	// Overlapping intervals
	fmt.Println(amountPainted([][]int{{1, 3}, {2, 5}, {3, 6}}))

	// Non-overlapping
	fmt.Println(amountPainted([][]int{{1, 2}, {3, 4}, {5, 6}}))
}

func amountPainted(paint [][]int) []int {
	maxEnd := 0
	for _, p := range paint {
		if p[1] > maxEnd {
			maxEnd = p[1]
		}
	}

	// DSU: next[x] = next unpainted point >= x
  // Alokasi slice integer
	next := make([]int, maxEnd+2)
  // Range loop: iterasi dengan indeks + nilai
	for i := range next {
		next[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if next[x] != x {
			next[x] = find(next[x])
		}
		return next[x]
	}

  // Alokasi slice integer
	result := make([]int, len(paint))
	for day, p := range paint {
		start, end := p[0], p[1]
		count := 0
		pos := find(start)
		for pos < end {
			count++
			next[pos] = find(pos + 1)
			pos = find(pos)
		}
		result[day] = count
	}

	return result
}
```
