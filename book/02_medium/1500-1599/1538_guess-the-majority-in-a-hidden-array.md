# 1538 — Guess The Majority In A Hidden Array

## Deskripsi

**Soal:** [1538. Guess The Majority In A Hidden Array](https://leetcode.com/problems/guess-the-majority-in-a-hidden-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1538: Guess the Majority in a Hidden Array
// https://leetcode.com/problems/guess-the-majority-in-a-hidden-array/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// Problem: There is an array of 0s and 1s. Use a query API that returns
	// whether the majority of 4 indices are same (0 or 1) or not.
	// Find the index of an element that is different from the majority,
	// or return -1 if all elements are the same.
	//
	// We simulate with a known array.

	arr := []int{0, 0, 1, 0, 0}
	idx := getMajorityIndex(len(arr), func(a, b, c, d int) bool {
		count0, count1 := 0, 0
		for _, i := range []int{a, b, c, d} {
			if arr[i] == 0 {
				count0++
			} else {
				count1++
			}
		}
		return count0 > count1 // returns true if majority is 0
	})
	fmt.Println("Different element index:", idx) // should be 2

	arr2 := []int{1, 1, 1, 1, 1}
	idx2 := getMajorityIndex(len(arr2), func(a, b, c, d int) bool {
		count0, count1 := 0, 0
		for _, i := range []int{a, b, c, d} {
			if arr2[i] == 0 {
				count0++
			} else {
				count1++
			}
		}
		return count0 > count1
	})
	fmt.Println("All same:", idx2) // should be -1
}

func getMajorityIndex(n int, query func(int, int, int, int) bool) int {
	// Time: O(N), Space: O(1)
	if n < 4 {
		return -1
	}

	// Compare 0,1,2,3 with 0,1,2,4 to see if 3 and 4 differ
	// If query(0,1,2,3) == query(0,1,2,4), then 3 and 4 are same
	// Otherwise they differ

	// Find two indices that differ
	diff := -1
	for i := 1; i < n; i++ {
		if query(0, 1, 2, i) != query(0, 1, 2, 0) {
			diff = i
			break
		}
	}

	if diff == -1 {
		return -1 // all elements are the same
	}

	// The differing element is either at position 0 or at diff
	// Check if 0,1,2,3 all agree
	if query(0, 1, 2, 0) == query(1, 0, 2, 0) {
		// They agree, so the majority is the common value, and
		// the minority is at diff
		return diff
	}
	return 0
}
```
