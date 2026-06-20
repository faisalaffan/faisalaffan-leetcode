# 0332 — Reconstruct Itinerary

## Deskripsi

**Soal:** [0332. Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func findItinerary(tickets [][]string) []string`

## Solusi Go

```go
package main

// LeetCode #332: Reconstruct Itinerary
// https://leetcode.com/problems/reconstruct-itinerary/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	// Build adjacency list with sorting
  // Membuat map untuk pencarian O(1): key → value
	graph := make(map[string][]string)
	for _, t := range tickets {
		from, to := t[0], t[1]
		graph[from] = append(graph[from], to)
	}
	for from := range graph {
		sort.Sort(sort.Reverse(sort.StringSlice(graph[from])))
	}

	// Hierholzer's algorithm: Eulerian path
	route := []string{}

	var dfs func(airport string)
	dfs = func(airport string) {
		for len(graph[airport]) > 0 {
			next := graph[airport][len(graph[airport])-1]
			graph[airport] = graph[airport][:len(graph[airport])-1]
			dfs(next)
		}
		route = append(route, airport)
	}

	dfs("JFK")

	// Reverse to get correct order
	for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
		route[i], route[j] = route[j], route[i]
	}
	return route
}

func main() {
	// Example 1
	tickets1 := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
		{"JFK", "ATL"},
		{"ATL", "JFK"},
	}
	fmt.Println(findItinerary(tickets1))
	// [JFK ATL JFK MUC LHR SFO SJC]

	// Example 2
	tickets2 := [][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}
	fmt.Println(findItinerary(tickets2))
	// [JFK ATL JFK SFO ATL SFO]

	// Example 3
	tickets3 := [][]string{
		{"JFK", "KUL"},
		{"JFK", "NRT"},
		{"NRT", "JFK"},
	}
	fmt.Println(findItinerary(tickets3))
	// [JFK NRT JFK KUL]
}
```
