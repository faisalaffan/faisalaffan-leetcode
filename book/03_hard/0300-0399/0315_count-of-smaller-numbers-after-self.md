# 0315 — Count Of Smaller Numbers After Self

## Deskripsi

**Soal:** [0315. Count Of Smaller Numbers After Self](https://leetcode.com/problems/count-of-smaller-numbers-after-self/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func countSmaller(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #315: Count of Smaller Numbers After Self
// https://leetcode.com/problems/count-of-smaller-numbers-after-self/
// Difficulty: Hard

import "fmt"

func countSmaller(nums []int) []int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return []int{}
	}

	// Pair each number with its original index
	type pair struct {
		val int
		idx int
	}
  // Membuat slice untuk menyimpan hasil
	arr := make([]pair, n)
	for i, v := range nums {
		arr[i] = pair{val: v, idx: i}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)

	var mergeSort func([]pair) []pair
	mergeSort = func(a []pair) []pair {
		if len(a) <= 1 {
			return a
		}
		mid := len(a) / 2
		left := mergeSort(a[:mid])
		right := mergeSort(a[mid:])

		// Merge while counting
  // Membuat slice untuk menyimpan hasil
		merged := make([]pair, 0, len(a))
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if left[i].val <= right[j].val {
				// All elements already placed from right that are smaller
				result[left[i].idx] += j
				merged = append(merged, left[i])
				i++
			} else {
				merged = append(merged, right[j])
				j++
			}
		}
		for i < len(left) {
			result[left[i].idx] += j
			merged = append(merged, left[i])
			i++
		}
		for j < len(right) {
			merged = append(merged, right[j])
			j++
		}
		return merged
	}

	mergeSort(arr)
	return result
}

func main() {
	// Example 1
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
	// [2, 1, 1, 0]

	// Example 2
	fmt.Println(countSmaller([]int{-1}))
	// [0]

	// Example 3
	fmt.Println(countSmaller([]int{-1, -1}))
	// [0, 0]
}
```
