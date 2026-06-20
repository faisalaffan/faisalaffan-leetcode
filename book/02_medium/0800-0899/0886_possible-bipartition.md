# 0886 — Possible Bipartition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func PossibleBipartition(n int, dislikes [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n + d) where d = len(dislikes)  
**Kompleksitas Ruang:** O(n + d)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #886: Possible Bipartition
// https://leetcode.com/problems/possible-bipartition/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PossibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	fmt.Println(PossibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(PossibleBipartition(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}))
}

// Time: O(n + d) where d = len(dislikes) | Space: O(n + d)
func PossibleBipartition(n int, dislikes [][]int) bool {
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n+1)
	for _, d := range dislikes {
		a, b := d[0], d[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

  // Alokasi slice integer
	color := make([]int, n+1) // 0 = uncolored, 1 = group A, -1 = group B

	var dfs func(node, c int) bool
	dfs = func(node, c int) bool {
		if color[node] != 0 {
			return color[node] == c
		}
		color[node] = c
		for _, nei := range graph[node] {
			if !dfs(nei, -c) {
				return false
			}
		}
		return true
	}

	for i := 1; i <= n; i++ {
		if color[i] == 0 && !dfs(i, 1) {
			return false
		}
	}

	return true
}
```
