package main

// LeetCode #1334: Find the City With the Smallest Number of Neighbors at a Threshold Distance
// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(findTheCity(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4)) // 3

	// Test case 2
	fmt.Println(findTheCity(5, [][]int{
		{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1},
	}, 2)) // 0

	// Test case 3
	fmt.Println(findTheCity(6, [][]int{
		{0, 3, 5}, {2, 3, 7}, {0, 5, 2}, {0, 2, 4}, {1, 4, 5}, {3, 4, 6},
	}, 10)) // 4
}

// Time: O(n^3) where n = number of cities (Floyd-Warshall)
// Space: O(n^2) for distance matrix
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i != j {
				dist[i][j] = math.MaxInt32
			}
		}
	}

	for _, e := range edges {
		dist[e[0]][e[1]] = e[2]
		dist[e[1]][e[0]] = e[2]
	}

	// Floyd-Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	minReachable := n
	result := -1
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				count++
			}
		}
		if count <= minReachable {
			minReachable = count
			result = i
		}
	}

	return result
}
