# 1743 — Restore The Array From Adjacent Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func restoreArray(adjacentPairs [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1743: Restore the Array From Adjacent Pairs
// https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func restoreArray(adjacentPairs [][]int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	graph := make(map[int][]int)
	for _, pair := range adjacentPairs {
		u, v := pair[0], pair[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Find the first element (has only 1 neighbor)
	start := 0
	for node, neighbors := range graph {
		if len(neighbors) == 1 {
			start = node
			break
		}
	}

	n := len(adjacentPairs) + 1
  // Alokasi slice integer
	result := make([]int, n)
	result[0] = start
	result[1] = graph[start][0]

	for i := 2; i < n; i++ {
		neighbors := graph[result[i-1]]
		if neighbors[0] == result[i-2] {
			result[i] = neighbors[1]
		} else {
			result[i] = neighbors[0]
		}
	}
	return result
}

func main() {
	fmt.Println(restoreArray([][]int{{2, 1}, {3, 4}, {3, 2}})) // Expected: [1, 2, 3, 4]
	fmt.Println(restoreArray([][]int{{4, -2}, {1, 4}, {-3, 1}})) // Expected: [-2, 4, 1, -3]
	fmt.Println(restoreArray([][]int{{100, -100}})) // Expected: [100, -100]
}
```
