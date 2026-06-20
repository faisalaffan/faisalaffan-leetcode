# 3237 — Alt And Tab Simulation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func simulationResult(windows []int, queries []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + q)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3237: Alt and Tab Simulation
// https://leetcode.com/problems/alt-and-tab-simulation/
// Difficulty: Medium [Paid]
// Time: O(n + q) | Space: O(n)

import (
	"container/list"
	"fmt"
)

func simulationResult(windows []int, queries []int) []int {
	order := list.New()
  // HashMap: O(1) lookup
	pos := make(map[int]*list.Element)
	for _, w := range windows {
		e := order.PushBack(w)
		pos[w] = e
	}

	for _, q := range queries {
		if e, ok := pos[q]; ok {
			order.MoveToFront(e)
		}
	}

  // Alokasi slice
	ans := make([]int, 0, order.Len())
	for e := order.Front(); e != nil; e = e.Next() {
		ans = append(ans, e.Value.(int))
	}
	return ans
}

func main() {
	fmt.Println(simulationResult([]int{1, 2, 3, 4}, []int{3, 1})) // Expected: [1 3 2 4]
	fmt.Println(simulationResult([]int{1, 2, 3}, []int{2}))        // Expected: [2 1 3]
}
```
