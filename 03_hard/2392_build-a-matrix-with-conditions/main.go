package main

// LeetCode #2392: Build a Matrix With Conditions
// https://leetcode.com/problems/build-a-matrix-with-conditions/
// Difficulty: Hard
//
// Given k (numbers 1..k), rowConditions (above/below pairs), and colConditions
// (left/right pairs), build a k x k matrix where each row and column has at most
// one number, and the conditions are satisfied. Numbers 1..k must each appear
// exactly once.
//
// Approach: Topological sort for rows using rowConditions, and for columns
// using colConditions. If either has a cycle, return empty matrix. Then place
// each number at (rowPos[number], colPos[number]).

import "fmt"

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowOrder := topoSort(k, rowConditions)
	colOrder := topoSort(k, colConditions)
	if rowOrder == nil || colOrder == nil {
		return [][]int{}
	}

	rowPos := make([]int, k+1)
	for i, v := range rowOrder {
		rowPos[v] = i
	}
	colPos := make([]int, k+1)
	for i, v := range colOrder {
		colPos[v] = i
	}

	matrix := make([][]int, k)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}
	for num := 1; num <= k; num++ {
		matrix[rowPos[num]][colPos[num]] = num
	}

	return matrix
}

func topoSort(k int, conditions [][]int) []int {
	graph := make([][]int, k+1)
	inDeg := make([]int, k+1)

	for _, c := range conditions {
		u, v := c[0], c[1]
		graph[u] = append(graph[u], v)
		inDeg[v]++
	}

	queue := make([]int, 0)
	for i := 1; i <= k; i++ {
		if inDeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := make([]int, 0, k)
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, v := range graph[u] {
			inDeg[v]--
			if inDeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(order) != k {
		return nil // cycle
	}
	return order
}

func main() {
	// Example 1
	k1 := 3
	rowC1 := [][]int{{1, 2}, {3, 2}}
	colC1 := [][]int{{2, 1}, {3, 2}}
	fmt.Println(buildMatrix(k1, rowC1, colC1))

	// Example 2: cycle
	k2 := 3
	rowC2 := [][]int{{1, 2}, {2, 3}, {3, 1}}
	colC2 := [][]int{{2, 3}}
	fmt.Println(buildMatrix(k2, rowC2, colC2))

	// Simple: k=2
	k3 := 2
	rowC3 := [][]int{{1, 2}}
	colC3 := [][]int{{2, 1}}
	fmt.Println(buildMatrix(k3, rowC3, colC3))

	// k=1
	k4 := 1
	fmt.Println(buildMatrix(k4, [][]int{}, [][]int{}))
}
