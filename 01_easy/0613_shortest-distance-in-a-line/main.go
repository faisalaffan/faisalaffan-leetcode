package main

// LeetCode #613: Shortest Distance in a Line
// https://leetcode.com/problems/shortest-distance-in-a-line/
// Difficulty: Easy [Paid]

import "fmt"

func ShortestDistanceInALine() string {
	return "SELECT MIN(ABS(p1.x - p2.x)) AS shortest FROM Point p1 JOIN Point p2 ON p1.x != p2.x"
}

func main() {
	fmt.Println(ShortestDistanceInALine())
}
