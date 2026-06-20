# 1462 — Course Schedule Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Floyd-Warshall

**Kompleksitas Waktu:** O(n^3) for Floyd-Warshall  
**Kompleksitas Ruang:** O(n^2) for reachability matrix

> **Untuk fresh graduate:** Kuasai dulu teknik **Floyd-Warshall** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1462: Course Schedule IV
// https://leetcode.com/problems/course-schedule-iv/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}))
	// [false, true]

	// Test case 2
	fmt.Println(checkIfPrerequisite(2, [][]int{}, [][]int{{1, 0}, {0, 1}}))
	// [false, false]

	// Test case 3
	fmt.Println(checkIfPrerequisite(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}}))
	// [true, false, true, false]
}

// Time: O(n^3) for Floyd-Warshall
// Space: O(n^2) for reachability matrix
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// Build adjacency list
  // Membuat matriks/slice 2D untuk DP
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
	}

	// Floyd-Warshall for reachability
  // Membuat matriks/slice 2D untuk DP
	reachable := make([][]bool, numCourses)
  // Range loop: iterasi dengan indeks + nilai
	for i := range reachable {
		reachable[i] = make([]bool, numCourses)
	}

	for _, p := range prerequisites {
		reachable[p[0]][p[1]] = true
	}

	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				if reachable[i][k] && reachable[k][j] {
					reachable[i][j] = true
				}
			}
		}
	}

	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = reachable[q[0]][q[1]]
	}

	return result
}
```
