# 1665 — Minimum Initial Energy To Finish Tasks

## Deskripsi

**Soal:** [1665. Minimum Initial Energy To Finish Tasks](https://leetcode.com/problems/minimum-initial-energy-to-finish-tasks/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1665: Minimum Initial Energy to Finish Tasks
// https://leetcode.com/problems/minimum-initial-energy-to-finish-tasks/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func main() {
	tasks1 := [][]int{{1, 2}, {2, 4}, {4, 8}}
	fmt.Printf("Test 1 - Input: [[1,2],[2,4],[4,8]]\nExpected: 8\nGot: %d\n\n", minimumEffort(tasks1))

	tasks2 := [][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}
	fmt.Printf("Test 2 - Input: [[1,3],[2,4],[10,11],[10,12],[8,9]]\nExpected: 32\nGot: %d\n\n", minimumEffort(tasks2))

	tasks3 := [][]int{{1, 1}, {1, 1}}
	fmt.Printf("Test 3 - Input: [[1,1],[1,1]]\nExpected: 1\nGot: %d\n", minimumEffort(tasks3))
}

func minimumEffort(tasks [][]int) int {
	// Sort by (minimum - actual) descending: tasks with the largest energy
	// deficit (minimum required vs actual consumed) should be done first.
	sort.Slice(tasks, func(i, j int) bool {
		return (tasks[i][1] - tasks[i][0]) > (tasks[j][1] - tasks[j][0])
	})

	result := 0
	curr := 0

	for _, t := range tasks {
		actual, minimum := t[0], t[1]
		if curr < minimum {
			result += minimum - curr
			curr = minimum
		}
		curr -= actual
	}

	return result
}
```
