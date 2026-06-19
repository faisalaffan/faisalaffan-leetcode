package main

// LeetCode #399: Evaluate Division
// https://leetcode.com/problems/evaluate-division/
// Difficulty: Medium
// Time: O(n + q*n) | Space: O(n)

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// Build graph
	graph := make(map[string]map[string]float64)
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]
		graph[b][a] = 1.0 / values[i]
	}

	var dfs func(src, dst string, visited map[string]bool) float64
	dfs = func(src, dst string, visited map[string]bool) float64 {
		if _, ok := graph[src]; !ok {
			return -1.0
		}
		if src == dst {
			return 1.0
		}
		visited[src] = true
		for neighbor, val := range graph[src] {
			if visited[neighbor] {
				continue
			}
			if neighbor == dst {
				return val
			}
			if result := dfs(neighbor, dst, visited); result != -1.0 {
				return val * result
			}
		}
		return -1.0
	}

	result := make([]float64, len(queries))
	for i, q := range queries {
		visited := make(map[string]bool)
		result[i] = dfs(q[0], q[1], visited)
	}
	return result
}

func main() {
	// Test case 1
	eq1 := [][]string{{"a", "b"}, {"b", "c"}}
	val1 := []float64{2.0, 3.0}
	q1 := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println("Test 1:", calcEquation(eq1, val1, q1))
	// Expected: [6.0, 0.5, -1.0, 1.0, -1.0]

	// Test case 2
	eq2 := [][]string{{"a", "b"}, {"c", "d"}}
	val2 := []float64{1.0, 1.0}
	q2 := [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}}
	fmt.Println("Test 2:", calcEquation(eq2, val2, q2))
	// Expected: [-1, -1, 1]
}
