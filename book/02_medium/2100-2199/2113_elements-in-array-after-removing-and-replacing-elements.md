# 2113 — Elements In Array After Removing And Replacing Elements

## Deskripsi

**Soal:** [2113. Elements In Array After Removing And Replacing Elements](https://leetcode.com/problems/elements-in-array-after-removing-and-replacing-elements/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func elementInArray(nums []int, queries [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #2113: Elements in Array After Removing and Replacing Elements
// https://leetcode.com/problems/elements-in-array-after-removing-and-replacing-elements/
// Difficulty: Medium [Paid]
// Time: O(n + q) | Space: O(1)

import "fmt"

func elementInArray(nums []int, queries [][]int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(queries))

	for i, q := range queries {
		time := q[0] % (2 * n)
		idx := q[1]
		result[i] = -1

		if time < n {
			// Removal phase: first 'time' elements removed
			// Remaining: nums[time:]
			if idx+time < n {
				result[i] = nums[idx+time]
			}
		} else if time > n {
			// Replacement phase: k = time-n (1..n-1)
			// Elements nums[n-k..n-1] placed back at front
			k := time - n
			if idx < k {
				result[i] = nums[n-k+idx]
			}
		}
		// time == n: array empty, result stays -1
	}
	return result
}

func main() {
	// Test case 1
	nums1 := []int{1, 2, 3}
	queries1 := [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {5, 1}}
	fmt.Println("Test 1:", elementInArray(nums1, queries1))
	// t=0: [1,2,3] → 1
	// t=1: [2,3] → 2
	// t=2: [3] → 3
	// t=3: [] → -1
	// t=4: [3] → 3
	// t=5: [2,3] → 2, 3
	// Expected: [1, 2, 3, -1, 3, 2, 3]

	// Test case 2
	nums2 := []int{5}
	queries2 := [][]int{{0, 0}, {1, 0}}
	fmt.Println("Test 2:", elementInArray(nums2, queries2))
	// t=0: [5] → 5
	// t=1: [] → -1
	// Expected: [5, -1]

	// Test case 3
	nums3 := []int{10, 20}
	queries3 := [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {0, 1}}
	fmt.Println("Test 3:", elementInArray(nums3, queries3))
	// Expected: [10, 20, -1, 20, 20]
}
```
