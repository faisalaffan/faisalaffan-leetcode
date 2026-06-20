# 0851 — Loud And Rich

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LoudAndRich(richer [][]int, quiet []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m) where m = len(richer)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #851: Loud and Rich
// https://leetcode.com/problems/loud-and-rich/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LoudAndRich([][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}))
	fmt.Println(LoudAndRich([][]int{{0, 1}, {1, 2}}, []int{0, 1, 2}))
	fmt.Println(LoudAndRich([][]int{}, []int{0}))
}

// Time: O(n + m) | Space: O(n + m) where m = len(richer)
func LoudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, n)
  // Alokasi slice integer
	indeg := make([]int, n)

	for _, r := range richer {
		a, b := r[0], r[1]
		graph[a] = append(graph[a], b)
		indeg[b]++
	}

  // Alokasi slice integer
	ans := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range ans {
		ans[i] = i
	}

	var queue []int
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range graph[u] {
			if quiet[ans[v]] > quiet[ans[u]] {
				ans[v] = ans[u]
			}
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return ans
}
```
