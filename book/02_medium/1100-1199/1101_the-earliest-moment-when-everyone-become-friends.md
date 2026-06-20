# 1101 — The Earliest Moment When Everyone Become Friends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func earliestAcq(logs [][]int, n int) int
```

> **💡 Hint:** Sort logs by timestamp, Union-Find

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** O(n log n + m * alpha(n)) where n = logs, m = N  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1101: The Earliest Moment When Everyone Become Friends
// https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friends/
// Difficulty: Medium
//
// Approach: Sort logs by timestamp, Union-Find
// Time: O(n log n + m * alpha(n)) where n = logs, m = N
// Space: O(N)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestAcq([][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}, 6)) // 20190301
	fmt.Println(earliestAcq([][]int{{0, 0, 1}, {1, 1, 2}, {2, 0, 2}}, 3)) // 2
}

func earliestAcq(logs [][]int, n int) int {
  // Custom sort dengan comparator
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

  // Alokasi slice integer
	parent := make([]int, n)
	for i := 0; i < n; i++ {
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
		pa, pb := find(a), find(b)
		if pa != pb {
			parent[pa] = pb
		}
	}

	groups := n
	for _, log := range logs {
		ts, a, b := log[0], log[1], log[2]
		if find(a) != find(b) {
			union(a, b)
			groups--
			if groups == 1 {
				return ts
			}
		}
	}

	return -1
}
```
