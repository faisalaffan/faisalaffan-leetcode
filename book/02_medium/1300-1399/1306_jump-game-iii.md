# 1306 — Jump Game Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func canReach(arr []int, start int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** O(n) where n is length of arr  |  **Ruang:** O(n) for visited array and queue

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1306: Jump Game III
// https://leetcode.com/problems/jump-game-iii/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5)) // true

	// Test case 2
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0)) // true

	// Test case 3
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2)) // false
}

// Time: O(n) where n is length of arr
// Space: O(n) for visited array and queue
func canReach(arr []int, start int) bool {
	n := len(arr)
	visited := make([]bool, n)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]

		if arr[i] == 0 {
			return true
		}

		for _, next := range []int{i + arr[i], i - arr[i]} {
			if next >= 0 && next < n && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return false
}
```
