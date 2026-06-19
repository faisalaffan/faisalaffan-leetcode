package main

// LeetCode #2092: Find All People With Secret
// https://leetcode.com/problems/find-all-people-with-secret/
// Difficulty: Hard
//
// Approach: Time-sorted Union-Find.
// Group meetings by time, union all participants within each time group,
// then check if any participant in the group is connected to a secret-knower.
// If not, reset their connections (they cannot learn the secret at this time).

import (
	"fmt"
	"sort"
)

func main() {
	// Example from problem statement
	n1 := 6
	meetings1 := [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}
	firstPerson1 := 1
	fmt.Printf("findAllPeople(%d, %v, %d) = %v (expected [0 1 2 3 5])\n",
		n1, meetings1, firstPerson1, findAllPeople(n1, meetings1, firstPerson1))

	// Additional tests
	n2 := 4
	meetings2 := [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}
	firstPerson2 := 3
	fmt.Printf("findAllPeople(%d, %v, %d) = %v\n",
		n2, meetings2, firstPerson2, findAllPeople(n2, meetings2, firstPerson2))

	n3 := 5
	meetings3 := [][]int{{0, 2, 1}, {1, 3, 1}, {4, 2, 2}}
	firstPerson3 := 2
	fmt.Printf("findAllPeople(%d, %v, %d) = %v\n",
		n3, meetings3, firstPerson3, findAllPeople(n3, meetings3, firstPerson3))
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// Sort meetings by time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	// Union-Find structure
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa != pb {
			parent[pa] = pb
		}
	}

	// 0 and firstPerson start with the secret
	union(0, firstPerson)

	i := 0
	for i < len(meetings) {
		t := meetings[i][2]
		j := i
		for j < len(meetings) && meetings[j][2] == t {
			j++
		}

		// Union all pairs in this time group
		for k := i; k < j; k++ {
			union(meetings[k][0], meetings[k][1])
		}

		// Collect unique people in this time group
		people := make(map[int]bool)
		for k := i; k < j; k++ {
			people[meetings[k][0]] = true
			people[meetings[k][1]] = true
		}

		// Reset those who are not connected to a secret-knower
		root0 := find(0)
		for p := range people {
			if find(p) != root0 {
				parent[p] = p
			}
		}

		i = j
	}

	// Collect all who know the secret
	result := []int{}
	root0 := find(0)
	for p := 0; p < n; p++ {
		if find(p) == root0 {
			result = append(result, p)
		}
	}
	return result
}
