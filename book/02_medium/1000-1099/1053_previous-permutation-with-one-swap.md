# 1053 — Previous Permutation With One Swap

## Deskripsi

**Soal:** [1053. Previous Permutation With One Swap](https://leetcode.com/problems/previous-permutation-with-one-swap/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

> **Ide Kunci:** Find rightmost pair where arr[i] > arr[i+1].

## Solusi Go

```go
package main

// LeetCode #1053: Previous Permutation With One Swap
// https://leetcode.com/problems/previous-permutation-with-one-swap/
// Difficulty: Medium
//
// Approach: Find rightmost pair where arr[i] > arr[i+1].
//           Swap with the largest smaller element to the right.
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))    // [3,1,2]
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))    // [1,1,5]
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7})) // [1,7,4,6,9]
}

func prevPermOpt1(arr []int) []int {
	i := len(arr) - 2
	for i >= 0 && arr[i] <= arr[i+1] {
		i--
	}

	if i < 0 {
		return arr
	}

	// Find rightmost smaller than arr[i]
	j := len(arr) - 1
	for arr[j] >= arr[i] {
		j--
	}
	// Skip duplicates
	for arr[j] == arr[j-1] {
		j--
	}

	arr[i], arr[j] = arr[j], arr[i]
	return arr
}
```
