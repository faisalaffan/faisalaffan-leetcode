# 0851 — Loud And Rich

## Deskripsi

**Soal:** [0851. Loud And Rich](https://leetcode.com/problems/loud-and-rich/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m) where m = len(richer)

**Algoritma:** —

## Solusi Go

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
  // Membuat slice 2D untuk DP/tabel
	graph := make([][]int, n)
  // Membuat slice untuk menyimpan hasil
	indeg := make([]int, n)

	for _, r := range richer {
		a, b := r[0], r[1]
		graph[a] = append(graph[a], b)
		indeg[b]++
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)
  // Iterasi seluruh elemen
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
