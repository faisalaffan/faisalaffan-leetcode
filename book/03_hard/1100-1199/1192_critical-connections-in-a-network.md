# 1192 — Critical Connections In A Network

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func criticalConnections(n int, connections [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1192: Critical Connections in a Network
// https://leetcode.com/problems/critical-connections-in-a-network/
// Difficulty: Hard

import "fmt"

func criticalConnections(n int, connections [][]int) [][]int {
	// Build adjacency list
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n)
	for _, e := range connections {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

  // Alokasi slice integer
	disc := make([]int, n) // discovery time
  // Alokasi slice integer
	low := make([]int, n)  // low-link value
  // Range loop: iterasi dengan indeks + nilai
	for i := range disc {
		disc[i] = -1
	}
	time := 0
	result := [][]int{}

	var dfs func(u, parent int)
	dfs = func(u, parent int) {
		disc[u] = time
		low[u] = time
		time++

		for _, v := range graph[u] {
			if v == parent {
				continue
			}
			if disc[v] == -1 {
				dfs(v, u)
				low[u] = min(low[u], low[v])
				// If low[v] > disc[u], edge (u,v) is a bridge
				if low[v] > disc[u] {
					result = append(result, []int{u, v})
				}
			} else {
				// Back edge
				low[u] = min(low[u], disc[v])
			}
		}
	}

	for i := 0; i < n; i++ {
		if disc[i] == -1 {
			dfs(i, -1)
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
	// Expected: [[1,3]]

	// Test case 2: star graph
	fmt.Println(criticalConnections(3, [][]int{{0, 1}, {0, 2}}))
	// Expected: [[0,1],[0,2]]

	// Test case 3: cycle
	fmt.Println(criticalConnections(3, [][]int{{0, 1}, {1, 2}, {2, 0}}))
	// Expected: []
}
```
