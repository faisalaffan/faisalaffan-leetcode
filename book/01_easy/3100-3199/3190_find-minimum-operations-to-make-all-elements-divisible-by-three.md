# 3190 — Find Minimum Operations To Make All Elements Divisible By Three

## Deskripsi

**Soal:** [3190. Find Minimum Operations To Make All Elements Divisible By Three](https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3190: Find Minimum Operations to Make All Elements Divisible by Three
// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMinimumOperationsToMakeAllElementsDivisibleByThree([]int{1, 2, 3, 4}))
	fmt.Println(FindMinimumOperationsToMakeAllElementsDivisibleByThree([]int{3, 6, 9}))
}

// FindMinimumOperationsToMakeAllElementsDivisibleByThree returns the minimum operations to make all elements divisible by 3.
// Each operation adds or subtracts 1 from an element.
// Time: O(n). Space: O(1).
func FindMinimumOperationsToMakeAllElementsDivisibleByThree(nums []int) int {
	ops := 0
	for _, num := range nums {
		r := num % 3
		if r == 1 || r == 2 {
			ops++
		}
	}
	return ops
}
```
