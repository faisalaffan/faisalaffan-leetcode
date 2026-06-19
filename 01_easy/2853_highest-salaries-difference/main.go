package main

// LeetCode #2853: Highest Salaries Difference
// https://leetcode.com/problems/highest-salaries-difference/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)
// Note: SQL problem, adapted to Go. Difference between highest salaries in two departments.

import "fmt"

func main() {
	salaries := []struct {
		Name       string
		Salary     int
		Department string
	}{
		{"Alice", 100000, "Engineering"},
		{"Bob", 90000, "Engineering"},
		{"Charlie", 80000, "Marketing"},
		{"David", 95000, "Marketing"},
	}
	fmt.Println(HighestSalariesDifference(salaries))
}

func HighestSalariesDifference(salaries []struct {
	Name       string
	Salary     int
	Department string
}) int {
	maxEng, maxMkt := 0, 0
	for _, s := range salaries {
		if s.Department == "Engineering" && s.Salary > maxEng {
			maxEng = s.Salary
		} else if s.Department == "Marketing" && s.Salary > maxMkt {
			maxMkt = s.Salary
		}
	}
	if maxEng > maxMkt {
		return maxEng - maxMkt
	}
	return maxMkt - maxEng
}
