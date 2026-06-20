# 0327 — Count Of Range Sum

## Deskripsi

**Soal:** [0327. Count Of Range Sum](https://leetcode.com/problems/count-of-range-sum/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Merge Sort (pengurutan gabung)

**Fungsi Solusi:** `func countRangeSum(nums []int, lower int, upper int) int`

## Solusi Go

```go
package main

// LeetCode #327: Count of Range Sum
// https://leetcode.com/problems/count-of-range-sum/
// Difficulty: Hard

import "fmt"

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int64, n+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + int64(v)
	}

	count := 0
	// Temporary buffer for merge sort
  // Membuat slice untuk menyimpan hasil
	temp := make([]int64, n+1)

	var mergeSort func(left, right int)
	mergeSort = func(left, right int) {
		if left >= right {
			return
		}
		mid := left + (right-left)/2
		mergeSort(left, mid)
		mergeSort(mid+1, right)

		// Count pairs crossing left and right halves
		// For each i in [left, mid], count j in [mid+1, right]
		// where lower <= prefix[j] - prefix[i] <= upper
		// i.e. prefix[i] + lower <= prefix[j] <= prefix[i] + upper
		lo, hi := mid+1, mid+1
		for i := left; i <= mid; i++ {
			for lo <= right && prefix[lo]-prefix[i] < int64(lower) {
				lo++
			}
			for hi <= right && prefix[hi]-prefix[i] <= int64(upper) {
				hi++
			}
			count += hi - lo
		}

		// Merge
		i, j, k := left, mid+1, left
		for i <= mid && j <= right {
			if prefix[i] <= prefix[j] {
				temp[k] = prefix[i]
				i++
			} else {
				temp[k] = prefix[j]
				j++
			}
			k++
		}
		for i <= mid {
			temp[k] = prefix[i]
			i++
			k++
		}
		for j <= right {
			temp[k] = prefix[j]
			j++
			k++
		}
		for i := left; i <= right; i++ {
			prefix[i] = temp[i]
		}
	}

	mergeSort(0, n)
	return count
}

func main() {
	// Example 1
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
	// 3

	// Example 2
	fmt.Println(countRangeSum([]int{0}, 0, 0))
	// 1

	// Example 3
	fmt.Println(countRangeSum([]int{2147483647, -2147483648, -1, 0}, -1, 0))
	// 4
}
```
