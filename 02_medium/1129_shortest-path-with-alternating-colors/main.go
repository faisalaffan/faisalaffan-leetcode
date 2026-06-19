package main

// LeetCode #1129: Shortest Path with Alternating Colors
// https://leetcode.com/problems/shortest-path-with-alternating-colors/
// Difficulty: Medium
//
// Approach: BFS with 2 states per node (reached by red edge / blue edge)
// Time: O(n + e)
// Space: O(n + e)

import "fmt"

func main() {
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))          // [0,1,-1]
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{1, 2}}))          // [0,1,2]
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	redGraph := make([][]int, n)
	blueGraph := make([][]int, n)

	for _, e := range redEdges {
		redGraph[e[0]] = append(redGraph[e[0]], e[1])
	}
	for _, e := range blueEdges {
		blueGraph[e[0]] = append(blueGraph[e[0]], e[1])
	}

	// dist[node][0] = distance reaching node via red edge
	// dist[node][1] = distance reaching node via blue edge
	dist := make([][2]int, n)
	for i := 1; i < n; i++ {
		dist[i] = [2]int{-1, -1}
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0}) // reached 0 via red (start counts as either)
	queue = append(queue, [2]int{0, 1}) // reached 0 via blue

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		node, color := cur[0], cur[1]
		nextColor := 1 - color

		if nextColor == 0 { // next edge should be red
			for _, next := range redGraph[node] {
				if dist[next][nextColor] == -1 {
					dist[next][nextColor] = dist[node][color] + 1
					queue = append(queue, [2]int{next, nextColor})
				}
			}
		} else { // next edge should be blue
			for _, next := range blueGraph[node] {
				if dist[next][nextColor] == -1 {
					dist[next][nextColor] = dist[node][color] + 1
					queue = append(queue, [2]int{next, nextColor})
				}
			}
		}
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = minDist(dist[i])
	}
	return result
}

func minDist(d [2]int) int {
	if d[0] == -1 {
		return d[1]
	}
	if d[1] == -1 {
		return d[0]
	}
	if d[0] < d[1] {
		return d[0]
	}
	return d[1]
}
