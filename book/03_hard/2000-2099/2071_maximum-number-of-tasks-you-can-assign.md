# 2071 — Maximum Number Of Tasks You Can Assign

## Deskripsi

**Soal:** [2071. Maximum Number Of Tasks You Can Assign](https://leetcode.com/problems/maximum-number-of-tasks-you-can-assign/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser), Binary Search (pencarian biner), Greedy (pemilihan optimal lokal)

**Fungsi Solusi:** `func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int`

> **Ide Kunci:** Binary Search + Multiset (simulated with sort + two-pointer)

## Solusi Go

```go
package main

// LeetCode #2071: Maximum Number of Tasks You Can Assign
// https://leetcode.com/problems/maximum-number-of-tasks-you-can-assign/
// Difficulty: Hard
// Approach: Binary Search + Multiset (simulated with sort + two-pointer)

import (
	"fmt"
	"sort"
)

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)

	// Binary search on how many tasks we can complete
	left, right := 0, len(tasks)
	if len(workers) < right {
		right = len(workers)
	}

  // Loop two-pointer: kiri vs kanan
	for left < right {
		mid := left + (right-left+1)/2 // try to do mid tasks (hardest mid tasks)
		if canAssign(tasks, workers, pills, strength, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func canAssign(tasks, workers []int, pills, strength, count int) bool {
	// We try to assign the 'count' hardest tasks to the 'count' strongest workers
	// tasks are sorted ascending, so we take the last 'count' tasks
	n := len(workers)
	taskIdx := len(tasks) - count // start from this task index

	// Use a multiset (multiset via sorted slice) for available workers
	// Actually, let's use a different approach: sort workers, then use a sliding window

	// Workers are sorted. We'll use a multiset implemented as a sorted slice
	// by maintaining available workers and removing with binary search.

	// Alternative: use two-pointer + multiset (e.g., a balanced BST via sorted slice + binary search)
	// We'll track workers that are available for the current task

	// Simpler approach: use a multiset implemented as a slice
  // Membuat slice untuk menyimpan hasil
	avail := make([]int, n)
	copy(avail, workers)
	// We'll pop from avail when a worker is used

	pillsLeft := pills

	// Process tasks from hardest to easiest (greedy: hardest task needs strongest worker)
	for j := len(tasks) - 1; j >= taskIdx; j-- {
		task := tasks[j]

		// Find a worker who can do this task without pill (weakest sufficient worker)
		idx := sort.SearchInts(avail, task)
		if idx < len(avail) {
			// Found a worker who can do it without pill, use them
			avail = append(avail[:idx], avail[idx+1:]...)
			continue
		}

		// Need a pill
		if pillsLeft > 0 {
			pillsLeft--
			// Find a worker who can do it with pill: worker + strength >= task
			needed := task - strength
			idx := sort.SearchInts(avail, needed)
			if idx < len(avail) {
				avail = append(avail[:idx], avail[idx+1:]...)
				continue
			}
		}

		// Cannot assign this task
		return false
	}

	return true
}

func main() {
	fmt.Println("2071. Maximum Number of Tasks You Can Assign")

	// Example 1
	tasks1 := []int{3, 2, 1}
	workers1 := []int{0, 3, 3}
	pills1 := 1
	strength1 := 1
	fmt.Printf("tasks=%v workers=%v pills=%d strength=%d → %d (expected 3)\n",
		tasks1, workers1, pills1, strength1, maxTaskAssign(tasks1, workers1, pills1, strength1))

	// Example 2
	tasks2 := []int{5, 4}
	workers2 := []int{0, 0, 0}
	pills2 := 1
	strength2 := 5
	fmt.Printf("tasks=%v workers=%v pills=%d strength=%d → %d (expected 1)\n",
		tasks2, workers2, pills2, strength2, maxTaskAssign(tasks2, workers2, pills2, strength2))

	// Example 3
	tasks3 := []int{10, 15, 30}
	workers3 := []int{0, 10, 20}
	pills3 := 2
	strength3 := 10
	fmt.Printf("tasks=%v workers=%v pills=%d strength=%d → %d (expected 2)\n",
		tasks3, workers3, pills3, strength3, maxTaskAssign(tasks3, workers3, pills3, strength3))
}
```
