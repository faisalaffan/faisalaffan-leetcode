# 2908 — Minimum Sum Of Mountain Triplets I

## Deskripsi

**Soal:** [2908. Minimum Sum Of Mountain Triplets I](https://leetcode.com/problems/minimum-sum-of-mountain-triplets-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2908: Minimum Sum of Mountain Triplets I
// https://leetcode.com/problems/minimum-sum-of-mountain-triplets-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumSum
	fmt.Println(MinimumSumOfMountainTripletsI([]int{8, 6, 1, 5, 3})) // 9
	fmt.Println(MinimumSumOfMountainTripletsI([]int{5, 4, 8, 7, 10, 2})) // 13
	fmt.Println(MinimumSumOfMountainTripletsI([]int{6, 5, 4, 3, 4, 5})) // -1
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: minimumSum
func MinimumSumOfMountainTripletsI(nums []int) int {
	n := len(nums)
	if n < 3 {
		return -1
	}

	// prefix[i] = min value in nums[0..i]
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < prefix[i-1] {
			prefix[i] = nums[i]
		} else {
			prefix[i] = prefix[i-1]
		}
	}

	// suffix[i] = min value in nums[i..n-1]
  // Membuat slice untuk menyimpan hasil
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffix[i+1] {
			suffix[i] = nums[i]
		} else {
			suffix[i] = suffix[i+1]
		}
	}

	minSum := -1
	for j := 1; j < n-1; j++ {
		if prefix[j-1] < nums[j] && suffix[j+1] < nums[j] {
			sum := prefix[j-1] + nums[j] + suffix[j+1]
			if minSum == -1 || sum < minSum {
				minSum = sum
			}
		}
	}
	return minSum
}
```
