# 3656 — Determine If A Simple Graph Exists

## Deskripsi

**Soal:** [3656. Determine If A Simple Graph Exists](https://leetcode.com/problems/determine-if-a-simple-graph-exists/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func determineIfASimpleGraphExists(degrees []int) bool`

## Solusi Go

```go
package main

// LeetCode #3656: Determine if a Simple Graph Exists
// https://leetcode.com/problems/determine-if-a-simple-graph-exists/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func determineIfASimpleGraphExists(degrees []int) bool {
	n := len(degrees)
  // Membuat slice untuk menyimpan hasil
	arr := make([]int, n)
	copy(arr, degrees)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			break
		}
		if arr[i] > n-i-1 {
			return false
		}
		for j := i + 1; j <= i+arr[i]; j++ {
			arr[j]--
			if arr[j] < 0 {
				return false
			}
		}
		arr[i] = 0
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	}

	return true
}

func main() {
	fmt.Println(determineIfASimpleGraphExists([]int{3, 3, 3, 3}))
	fmt.Println(determineIfASimpleGraphExists([]int{1, 1, 0}))
	fmt.Println(determineIfASimpleGraphExists([]int{1, 1, 1}))
}
```
