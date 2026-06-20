# 2860 — Happy Students

## Deskripsi

**Soal:** [2860. Happy Students](https://leetcode.com/problems/happy-students/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func HappyStudents(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2860: Happy Students
// https://leetcode.com/problems/happy-students/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func HappyStudents(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	count := 0

	// Check for empty selection (all students are happy)
	if nums[0] != 0 {
		count++
	}

	for i := 0; i < n; i++ {
		selected := i + 1
		// nums[i] < selected: the student at position i needs fewer than selected
		if nums[i] < selected {
			// Check if next student (if exists) needs more than selected
			if i+1 >= n || nums[i+1] > selected {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(HappyStudents([]int{1, 1}))
	fmt.Println(HappyStudents([]int{6, 0, 3, 3, 6, 7, 2, 7}))
}
```
