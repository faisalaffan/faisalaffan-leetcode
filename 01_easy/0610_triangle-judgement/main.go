package main

// LeetCode #610: Triangle Judgement
// https://leetcode.com/problems/triangle-judgement/
// Difficulty: Easy

import "fmt"

func TriangleJudgement() string {
	return "SELECT x, y, z, CASE WHEN x + y > z AND x + z > y AND y + z > x THEN 'Yes' ELSE 'No' END AS triangle FROM Triangle"
}

func main() {
	fmt.Println(TriangleJudgement())
}
