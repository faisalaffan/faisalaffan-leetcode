# 3554 — Find Category Recommendation Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func findCategoryPairs(categories []string, relations [][]string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3554: Find Category Recommendation Pairs
// https://leetcode.com/problems/find-category-recommendation-pairs/
// Difficulty: Hard
//
// Given categories and relationships, find all recommendation pairs.
// This is originally a SQL problem. Implement as Go function.
//
// Approach: Build a graph of category relationships, find pairs that
// meet the recommendation criteria.

import "fmt"
import "sort"

func main() {
	// Example 1
	fmt.Println(findCategoryPairs([]string{"A", "B", "C"}, [][]string{{"A", "B"}, {"B", "C"}}))
	// Example 2: no relationships
	fmt.Println(findCategoryPairs([]string{"X", "Y"}, [][]string{}))
	// Edge: single category
	fmt.Println(findCategoryPairs([]string{"A"}, [][]string{}))
}

func findCategoryPairs(categories []string, relations [][]string) [][]string {
	// Build adjacency
  // HashMap: O(1) lookup
	adj := make(map[string]map[string]bool)
	for _, cat := range categories {
		adj[cat] = make(map[string]bool)
	}
	for _, rel := range relations {
		a, b := rel[0], rel[1]
		adj[a][b] = true
		adj[b][a] = true
	}

	// Find all recommendation pairs (mutual friends)
	type pair struct {
		a, b string
	}
	var pairs []pair
  // HashMap: O(1) lookup
	seen := make(map[string]bool)

	for _, a := range categories {
		for _, b := range categories {
			if a >= b {
				continue
			}
			key := a + ":" + b
			if seen[key] {
				continue
			}
			// Check if they share a common connection
			for _, c := range categories {
				if c == a || c == b {
					continue
				}
				if adj[a][c] && adj[b][c] {
					pairs = append(pairs, pair{a, b})
					seen[key] = true
					break
				}
			}
		}
	}

  // Custom sort
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].a != pairs[j].a {
			return pairs[i].a < pairs[j].a
		}
		return pairs[i].b < pairs[j].b
	})

  // Matriks 2D
	result := make([][]string, len(pairs))
	for i, p := range pairs {
		result[i] = []string{p.a, p.b}
	}
	return result
}
```
