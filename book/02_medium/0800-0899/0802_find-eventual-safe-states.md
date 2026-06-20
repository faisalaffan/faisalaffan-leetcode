# 0802 — Find Eventual Safe States

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func eventualSafeNodes(graph [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(V + E)  
**Kompleksitas Ruang:** O(V)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #802: Find Eventual Safe States
// https://leetcode.com/problems/find-eventual-safe-states/
// Difficulty: Medium
// Time: O(V + E)
// Space: O(V)

import "fmt"

func main() {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
	fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
}

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
  // Alokasi slice integer
	state := make([]int, n) // 0=unvisited, 1=visiting, 2=safe

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if state[node] > 0 {
			return state[node] == 2
		}

		state[node] = 1
		for _, neighbor := range graph[node] {
			if !dfs(neighbor) {
				return false
			}
		}
		state[node] = 2
		return true
	}

  // Alokasi slice integer
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if dfs(i) {
			result = append(result, i)
		}
	}

	return result
}
```
