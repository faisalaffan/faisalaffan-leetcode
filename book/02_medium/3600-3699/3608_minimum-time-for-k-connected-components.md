# 3608 — Minimum Time For K Connected Components

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumTimeForKConnectedComponents(n int, edges [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3608: Minimum Time for K Connected Components
// https://leetcode.com/problems/minimum-time-for-k-connected-components/
// Difficulty: Medium
// Complexity: O(n + m) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	n := 4
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}
	k := 2
	fmt.Println("Test 1:", MinimumTimeForKConnectedComponents(n, edges, k))
	// Test case 2
	n2 := 5
	edges2 := [][]int{{0, 1}, {2, 3}}
	k2 := 3
	fmt.Println("Test 2:", MinimumTimeForKConnectedComponents(n2, edges2, k2))
	// Test case 3
	n3 := 3
	edges3 := [][]int{{0, 1}, {1, 2}, {0, 2}}
	k3 := 1
	fmt.Println("Test 3:", MinimumTimeForKConnectedComponents(n3, edges3, k3))
}

func MinimumTimeForKConnectedComponents(n int, edges [][]int, k int) int {
	// Find connected components count
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	components := 0
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			components++
			dfs(i)
		}
	}
	if components >= k {
		return 0
	}
	return k - components
}
```
