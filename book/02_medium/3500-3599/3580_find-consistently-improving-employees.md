# 3580 — Find Consistently Improving Employees

## Deskripsi

**Soal:** [3580. Find Consistently Improving Employees](https://leetcode.com/problems/find-consistently-improving-employees/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3580: Find Consistently Improving Employees
// https://leetcode.com/problems/find-consistently-improving-employees/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	scores := []int{1, 2, 3, 4, 5}
	fmt.Println("Test 1:", FindConsistentlyImprovingEmployees(scores))
	// Test case 2
	scores2 := []int{5, 4, 3, 2, 1}
	fmt.Println("Test 2:", FindConsistentlyImprovingEmployees(scores2))
	// Test case 3
	scores3 := []int{1, 3, 2, 4, 5}
	fmt.Println("Test 3:", FindConsistentlyImprovingEmployees(scores3))
}

func FindConsistentlyImprovingEmployees(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	count := 0
	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[i-1] {
			count++
		}
	}
	return count
}
```
