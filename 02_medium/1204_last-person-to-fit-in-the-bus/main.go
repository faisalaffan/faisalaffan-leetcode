package main

import (
	"fmt"
	"sort"
)

// LeetCode #1204: Last Person to Fit in the Bus
// https://leetcode.com/problems/last-person-to-fit-in-the-bus/
// Difficulty: Medium

// People queue for bus with weight limit 1000.
// Find the last person name that can board without exceeding limit.

// Time: O(n log n)
// Space: O(n)

type person struct {
	name   string
	weight int
	turn   int
}

func lastToFit(people []person) string {
	sort.Slice(people, func(i, j int) bool {
		return people[i].turn < people[j].turn
	})

	total := 0
	lastName := ""
	for _, p := range people {
		if total+p.weight <= 1000 {
			total += p.weight
			lastName = p.name
		} else {
			break
		}
	}
	return lastName
}

func main() {
	people := []person{
		{"Alice", 200, 1},
		{"Bob", 300, 2},
		{"Charlie", 400, 3},
		{"Dave", 200, 4},
	}
	fmt.Printf("%q (expected: \"Charlie\")\n", lastToFit(people))

	people2 := []person{
		{"John", 500, 1},
		{"Jane", 600, 2},
	}
	fmt.Printf("%q (expected: \"John\")\n", lastToFit(people2))
}
