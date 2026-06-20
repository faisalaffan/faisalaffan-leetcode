# 0582 — Kill Process

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func KillProcess(pid []int, ppid []int, kill int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #582: Kill Process
// https://leetcode.com/problems/kill-process/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}
	kill := 5
	fmt.Println(KillProcess(pid, ppid, kill))
}

func KillProcess(pid []int, ppid []int, kill int) []int {
	// Build adjacency list: parent -> children
  // HashMap: O(1) lookup
	children := make(map[int][]int)
	for i, p := range ppid {
		children[p] = append(children[p], pid[i])
	}

	// BFS/DFS to find all processes to kill
	result := []int{}
	queue := []int{kill}
	for len(queue) > 0 {
		process := queue[0]
		queue = queue[1:]
		result = append(result, process)
		queue = append(queue, children[process]...)
	}

	return result
}
```
