# 1203 — Sort Items By Groups Respecting Dependencies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func sortItems(n int, m int, group []int, beforeItems [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, BFS, Topological Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1203: Sort Items by Groups Respecting Dependencies
// https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/
// Difficulty: Hard
//
// There are n items, each belonging to a group (or -1 for no group). There are
// also dependency lists: beforeItems[i] contains items that must come before
// item i. Items within the same group must be adjacent in the final order.
// Return any valid ordering, or an empty slice if impossible.

import "fmt"

func main() {
	// Example 1
	n := 8
	group := []int{-1, -1, 1, 0, 0, 1, 0, -1}
	beforeItems := [][]int{
		{}, {6}, {5}, {6}, {3, 6}, {}, {}, {},
	}
	fmt.Println(sortItems(n, 0, group, beforeItems))

	// Example 2 (impossible due to cycle)
	n2 := 8
	group2 := []int{-1, -1, 1, 0, 0, 1, 0, -1}
	beforeItems2 := [][]int{
		{}, {6}, {5}, {6}, {3}, {}, {4}, {},
	}
	fmt.Println(sortItems(n2, 0, group2, beforeItems2))

	// Simple case
	n3 := 3
	group3 := []int{0, 0, 0}
	beforeItems3 := [][]int{
		{}, {0}, {1},
	}
	fmt.Println(sortItems(n3, 0, group3, beforeItems3))
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// Assign unique group IDs to items with group == -1
	nextGroup := m
	for i := 0; i < n; i++ {
		if group[i] == -1 {
			group[i] = nextGroup
			nextGroup++
		}
	}
	numGroups := nextGroup

	// --- Item-level topological sort ---
  // Membuat matriks/slice 2D untuk DP
	itemGraph := make([][]int, n)
  // Alokasi slice integer
	itemInDeg := make([]int, n)
	for i := 0; i < n; i++ {
		for _, dep := range beforeItems[i] {
			// Only add dependency if items are in different groups,
			// OR if there is no group change (same group).
			// Actually, we always add item-level deps; group adjacency
			// is handled via group-level sorting.
			itemGraph[dep] = append(itemGraph[dep], i)
			itemInDeg[i]++
		}
	}

	itemOrder := topologicalSort(n, itemGraph, itemInDeg)
	if len(itemOrder) == 0 {
		return []int{}
	}

	// --- Group-level topological sort ---
  // Membuat matriks/slice 2D untuk DP
	groupGraph := make([][]int, numGroups)
  // Alokasi slice integer
	groupInDeg := make([]int, numGroups)

	for i := 0; i < n; i++ {
		for _, dep := range beforeItems[i] {
			gFrom := group[dep]
			gTo := group[i]
			if gFrom != gTo {
				groupGraph[gFrom] = append(groupGraph[gFrom], gTo)
			}
		}
	}

	// Deduplicate group edges (otherwise in-degree may be inflated)
	for g := 0; g < numGroups; g++ {
  // Membuat map (HashMap) — pencarian O(1)
		seen := make(map[int]bool)
  // Alokasi slice integer
		uniq := make([]int, 0, len(groupGraph[g]))
		for _, to := range groupGraph[g] {
			if !seen[to] {
				seen[to] = true
				uniq = append(uniq, to)
				groupInDeg[to]++
			}
		}
		groupGraph[g] = uniq
	}

	groupOrder := topologicalSort(numGroups, groupGraph, groupInDeg)
	if len(groupOrder) == 0 {
		return []int{}
	}

	// Group items by their group order
  // Membuat map (HashMap) — pencarian O(1)
	groupItems := make(map[int][]int) // group -> items in itemOrder
	for _, it := range itemOrder {
		g := group[it]
		groupItems[g] = append(groupItems[g], it)
	}

	// Concatenate in group order
  // Alokasi slice integer
	result := make([]int, 0, n)
	for _, g := range groupOrder {
		result = append(result, groupItems[g]...)
	}

	return result
}

// topologicalSort performs Kahn's algorithm. Returns empty slice if a cycle exists.
func topologicalSort(n int, graph [][]int, inDeg []int) []int {
  // Alokasi slice integer
	inDegCopy := make([]int, n)
	copy(inDegCopy, inDeg)

  // Alokasi slice integer
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if inDegCopy[i] == 0 {
			queue = append(queue, i)
		}
	}

  // Alokasi slice integer
	result := make([]int, 0, n)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			inDegCopy[neighbor]--
			if inDegCopy[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) < n {
		return []int{} // cycle
	}
	return result
}
```
