# 0444 — Sequence Reconstruction

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sequenceReconstruction(org []int, seqs [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #444: Sequence Reconstruction
// https://leetcode.com/problems/sequence-reconstruction/
// Difficulty: Medium [Paid]
// Time: O(n + m) | Space: O(n)

import "fmt"

func sequenceReconstruction(org []int, seqs [][]int) bool {
	n := len(org)
	// Build indegree map and edges
  // Alokasi slice integer
	indegree := make([]int, n+1)
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n+1)
	exists := make([]bool, n+1)

	for _, seq := range seqs {
		for _, num := range seq {
			if num < 1 || num > n {
				return false
			}
			exists[num] = true
		}
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(seq)-1; i++ {
			u, v := seq[i], seq[i+1]
			graph[u] = append(graph[u], v)
			indegree[v]++
		}
	}

	// Check all numbers exist
	for i := 1; i <= n; i++ {
		if !exists[i] {
			return false
		}
	}

	// BFS: only one node with indegree 0 at each step
	queue := []int{}
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	idx := 0
	for len(queue) == 1 {
		u := queue[0]
		queue = queue[1:]
		if u != org[idx] {
			return false
		}
		idx++
		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return idx == n
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	// Expected: false (cycle)
}
```
