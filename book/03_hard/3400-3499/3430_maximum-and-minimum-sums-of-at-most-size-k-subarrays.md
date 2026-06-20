# 3430 — Maximum And Minimum Sums Of At Most Size K Subarrays

## Deskripsi

**Soal:** [3430. Maximum And Minimum Sums Of At Most Size K Subarrays](https://leetcode.com/problems/maximum-and-minimum-sums-of-at-most-size-k-subarrays/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO), Monotonic Stack (tumpukan monoton)

**Fungsi Solusi:** `func min(a, b int) int`

## Solusi Go

```go
package main

// LeetCode #3430: Maximum and Minimum Sums of at Most Size K Subarrays
// https://leetcode.com/problems/maximum-and-minimum-sums-of-at-most-size-k-subarrays/
// Difficulty: Hard
//
// Monotonic stack for previous/next smaller/greater element.
// Contribution counting with subarray length constraint ≤ k.

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// For each element, compute how many subarrays of length ≤ k it is min/max of.
func countContributions(leftDist, rightDist, k int) int64 {
	// leftDist: distance to previous smaller/greater element
	// rightDist: distance to next smaller/greater element
	// Number of subarrays of length ≤ k where this element is min (or max):
	// choose left extension a in [1, leftDist], right extension b in [1, rightDist]
	// subarray length = a + b - 1 ≤ k  =>  a + b ≤ k + 1
	var total int64
	maxA := min(leftDist, k)
	for a := 1; a <= maxA; a++ {
		bLimit := min(rightDist, k+1-a)
		if bLimit > 0 {
			total += int64(bLimit)
		}
	}
	return total
}

func minMaxSumOfSubarraysAtMostK(nums []int, k int) (int64, int64) {
	n := len(nums)

	// --- Monotonic stacks ---

	// Previous smaller (strict)
  // Membuat slice untuk menyimpan hasil
	ps := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ps[i] = stack[len(stack)-1]
		} else {
			ps[i] = -1
		}
		stack = append(stack, i)
	}

	// Next smaller (strict: nums[ns[i]] < nums[i])
  // Membuat slice untuk menyimpan hasil
	ns := make([]int, n)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ns[i] = stack[len(stack)-1]
		} else {
			ns[i] = n
		}
		stack = append(stack, i)
	}

	// Previous greater (strict)
  // Membuat slice untuk menyimpan hasil
	pg := make([]int, n)
	stack = stack[:0]
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			pg[i] = stack[len(stack)-1]
		} else {
			pg[i] = -1
		}
		stack = append(stack, i)
	}

	// Next greater (strict: nums[ng[i]] > nums[i])
  // Membuat slice untuk menyimpan hasil
	ng := make([]int, n)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ng[i] = stack[len(stack)-1]
		} else {
			ng[i] = n
		}
		stack = append(stack, i)
	}

	var sumMin, sumMax int64
	for i := 0; i < n; i++ {
		lMin := i - ps[i]    // left extension limit for min contribution
		rMin := ns[i] - i    // right extension limit for min contribution
		lMax := i - pg[i]    // left extension limit for max contribution
		rMax := ng[i] - i    // right extension limit for max contribution

		cntMin := countContributions(lMin, rMin, k)
		cntMax := countContributions(lMax, rMax, k)

		sumMin += int64(nums[i]) * cntMin
		sumMax += int64(nums[i]) * cntMax
	}
	return sumMin, sumMax
}

func main() {
	// Example test: [1,2,3,4,5], k=3
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	minS, maxS := minMaxSumOfSubarraysAtMostK(nums, k)
	fmt.Printf("nums=%v, k=%d -> minSum=%d, maxSum=%d\n", nums, k, minS, maxS)

	// Test: [1,3,2], k=2
	nums2 := []int{1, 3, 2}
	k2 := 2
	minS2, maxS2 := minMaxSumOfSubarraysAtMostK(nums2, k2)
	fmt.Printf("nums=%v, k=%d -> minSum=%d, maxSum=%d\n", nums2, k2, minS2, maxS2)

	// Test: single element
	nums3 := []int{5}
	k3 := 1
	minS3, maxS3 := minMaxSumOfSubarraysAtMostK(nums3, k3)
	fmt.Printf("nums=%v, k=%d -> minSum=%d, maxSum=%d\n", nums3, k3, minS3, maxS3)

	// Test: all equal
	nums4 := []int{2, 2, 2, 2}
	k4 := 2
	minS4, maxS4 := minMaxSumOfSubarraysAtMostK(nums4, k4)
	fmt.Printf("nums=%v, k=%d -> minSum=%d, maxSum=%d\n", nums4, k4, minS4, maxS4)
}
```
