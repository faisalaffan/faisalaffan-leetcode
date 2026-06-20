# 0207 — Course Schedule

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func canFinish(numCourses int, prerequisites [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(V+E), Space: O(V+E)  
**Kompleksitas Ruang:** O(V+E)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #207: Course Schedule
// https://leetcode.com/problems/course-schedule/
// Difficulty: Medium
// Time: O(V+E), Space: O(V+E)

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
  // Membuat matriks/slice 2D untuk DP
	graph := make([][]int, numCourses)
  // Alokasi slice integer
	inDegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		graph[prereq] = append(graph[prereq], course)
		inDegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return count == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {2, 1}, {3, 2}}))
}
```
