# 2936 — Number Of Equal Numbers Blocks

## Deskripsi

**Soal:** [2936. Number Of Equal Numbers Blocks](https://leetcode.com/problems/number-of-equal-numbers-blocks/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(k log n)  
**Kompleksitas Ruang:** O(1) where k = number of blocks

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func countBlocks(arr *BigArray) int`

## Solusi Go

```go
package main

// LeetCode #2936: Number of Equal Numbers Blocks
// https://leetcode.com/problems/number-of-equal-numbers-blocks/
// Difficulty: Medium

import "fmt"

type BigArray struct {
	data []int
}

func (b *BigArray) at(index int) int {
	return b.data[index]
}

func (b *BigArray) size() int {
	return len(b.data)
}

func countBlocks(arr *BigArray) int {
	n := arr.size()
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	blocks := 0
	i := 0

	for i < n {
		blocks++
		val := arr.at(i)

		// Binary search to find the rightmost position where arr.at(pos) == val
		left, right := i, n-1
  // Loop two-pointer: kiri vs kanan
		for left < right {
			mid := (left + right + 1) / 2
			if arr.at(mid) == val {
				left = mid
			} else {
				right = mid - 1
			}
		}

		i = left + 1
	}

	return blocks
}

func main() {
	// Test case 1: [1,1,1,2,2,3,3,3,3] -> 3 blocks
	arr1 := &BigArray{data: []int{1, 1, 1, 2, 2, 3, 3, 3, 3}}
	fmt.Println(countBlocks(arr1)) // 3

	// Test case 2: [5,5,5,5,5] -> 1 block
	arr2 := &BigArray{data: []int{5, 5, 5, 5, 5}}
	fmt.Println(countBlocks(arr2)) // 1

	// Test case 3: [1,2,3,4,5] -> 5 blocks
	arr3 := &BigArray{data: []int{1, 2, 3, 4, 5}}
	fmt.Println(countBlocks(arr3)) // 5

	// Test case 4: empty array -> 0 blocks
	arr4 := &BigArray{data: []int{}}
	fmt.Println(countBlocks(arr4)) // 0

	// Test case 5: single element -> 1 block
	arr5 := &BigArray{data: []int{42}}
	fmt.Println(countBlocks(arr5)) // 1
}

// Time: O(k log n) | Space: O(1) where k = number of blocks
```
