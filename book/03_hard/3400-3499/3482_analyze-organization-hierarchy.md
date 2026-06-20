# 3482 — Analyze Organization Hierarchy

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func analyzeOrg(org [][]string) []Employee
```

> **💡 Hint:** Build the org tree from edges and compute depth/level

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3482: Analyze Organization Hierarchy
// https://leetcode.com/problems/analyze-organization-hierarchy/
// Difficulty: Hard
//
// Given an organization hierarchy represented as employee-manager pairs,
// analyze the hierarchy. This is originally a SQL problem.
//
// Approach: Build the org tree from edges and compute depth/level
// for each employee using BFS/DFS from root.

import (
	"fmt"
	"sort"
)

func main() {
	// Example
	fmt.Println(analyzeOrg([][]string{{"Alice", "Bob"}, {"Bob", "Charlie"}, {"Charlie", ""}}))
	// Flat org
	fmt.Println(analyzeOrg([][]string{{"Alice", ""}, {"Bob", ""}, {"Charlie", ""}}))
	// Deep hierarchy
	fmt.Println(analyzeOrg([][]string{{"E1", "E2"}, {"E2", "E3"}, {"E3", "E4"}, {"E4", ""}}))
}

type Employee struct {
	Name  string
	Depth int
}

func analyzeOrg(org [][]string) []Employee {
	// Build parent -> children map
  // Membuat map (HashMap) — pencarian O(1)
	children := make(map[string][]string)
  // Membuat map (HashMap) — pencarian O(1)
	parent := make(map[string]string)
  // Membuat map (HashMap) — pencarian O(1)
	allEmps := make(map[string]bool)

	for _, rel := range org {
		emp, mgr := rel[0], rel[1]
		allEmps[emp] = true
		if mgr != "" {
			allEmps[mgr] = true
			children[mgr] = append(children[mgr], emp)
			parent[emp] = mgr
		}
	}

	// Find root(s) - employees with no manager
	var roots []string
	for emp := range allEmps {
		if _, ok := parent[emp]; !ok {
			roots = append(roots, emp)
		}
	}
	sort.Strings(roots)

	// BFS from each root to compute depth
	var result []Employee
	for _, root := range roots {
		queue := []struct {
			name  string
			depth int
		}{{root, 0}}

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			result = append(result, Employee{curr.name, curr.depth})
			for _, child := range children[curr.name] {
				queue = append(queue, struct {
					name  string
					depth int
				}{child, curr.depth + 1})
			}
		}
	}

  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		if result[i].Depth != result[j].Depth {
			return result[i].Depth < result[j].Depth
		}
		return result[i].Name < result[j].Name
	})

	return result
}
```
