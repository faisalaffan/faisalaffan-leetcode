# 0645 — Set Mismatch

## Deskripsi

**Soal:** [0645. Set Mismatch](https://leetcode.com/problems/set-mismatch/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #645: Set Mismatch
// https://leetcode.com/problems/set-mismatch/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4})) // [2, 3]
	fmt.Println(findErrorNums([]int{1, 1}))        // [1, 2]
	fmt.Println(findErrorNums([]int{2, 2}))        // [2, 1]
}

// findErrorNums finds the duplicated and missing number in the set.
// Time: O(n). Space: O(1).
func findErrorNums(nums []int) []int {
	n := len(nums)
	sum := 0
	sumSq := 0
	expectedSum := n * (n + 1) / 2
	expectedSumSq := n * (n + 1) * (2*n + 1) / 6

	for _, v := range nums {
		sum += v
		sumSq += v * v
	}

	// diff = duplicate - missing
	diff := sum - expectedSum
	// sqDiff = duplicate^2 - missing^2
	sqDiff := sumSq - expectedSumSq
	// duplicate + missing = sqDiff / diff
	plus := sqDiff / diff

	dup := (diff + plus) / 2
	miss := (plus - diff) / 2
	return []int{dup, miss}
}
```
