# 3432 — Count Partitions With Even Sum Difference

## Deskripsi

**Soal:** [3432. Count Partitions With Even Sum Difference](https://leetcode.com/problems/count-partitions-with-even-sum-difference/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3432: Count Partitions with Even Sum Difference
// https://leetcode.com/problems/count-partitions-with-even-sum-difference/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(CountPartitionsWithEvenSumDifference([]int{10, 10, 10, 10, 10}))
}

// CountPartitionsWithEvenSumDifference counts partitions where the difference between left and right sums is even.
// Time: O(n). Space: O(1).
func CountPartitionsWithEvenSumDifference(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	leftSum := 0
	count := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		rightSum := totalSum - leftSum
		if (leftSum-rightSum)%2 == 0 {
			count++
		}
	}
	return count
}
```
