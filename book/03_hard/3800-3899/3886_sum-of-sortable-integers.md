# 3886 — Sum Of Sortable Integers

## Deskripsi

**Soal:** [3886. Sum Of Sortable Integers](https://leetcode.com/problems/sum-of-sortable-integers/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3886: Sum of Sortable Integers
// https://leetcode.com/problems/sum-of-sortable-integers/
// Difficulty: Hard
//
// An integer is "sortable" if its digits can be rearranged to form
// a strictly increasing sequence (i.e., no adjacent equal digits
// after sorting). Equivalently, the multiset of digits must have
// no duplicates.
//
// Return the sum of all sortable integers in the given array.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumOfSortableIntegers([]int{12, 21, 11}))
	// Example 2
	fmt.Println(sumOfSortableIntegers([]int{123, 111, 122}))
	// Edge: single digit
	fmt.Println(sumOfSortableIntegers([]int{5, 7, 9}))
	// Edge: all duplicates
	fmt.Println(sumOfSortableIntegers([]int{11, 22, 33}))
}

func sumOfSortableIntegers(nums []int) int {
	sum := 0
	for _, num := range nums {
		if isSortable(num) {
			sum += num
		}
	}
	return sum
}

func isSortable(n int) bool {
  // Edge case: input kosong
	if n == 0 {
		return true
	}
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)
	for n > 0 {
		d := n % 10
		if seen[d] {
			return false
		}
		seen[d] = true
		n /= 10
	}
	return true
}
```
