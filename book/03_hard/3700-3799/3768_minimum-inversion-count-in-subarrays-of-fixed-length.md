# 3768 — Minimum Inversion Count In Subarrays Of Fixed Length

## Deskripsi

**Soal:** [3768. Minimum Inversion Count In Subarrays Of Fixed Length](https://leetcode.com/problems/minimum-inversion-count-in-subarrays-of-fixed-length/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser), Fenwick Tree (Binary Indexed Tree)

> **Ide Kunci:** Sliding window with Fenwick tree. Maintain inversion

## Solusi Go

```go
package main

// LeetCode #3768: Minimum Inversion Count in Subarrays of Fixed Length
// https://leetcode.com/problems/minimum-inversion-count-in-subarrays-of-fixed-length/
// Difficulty: Hard
//
// Find min inversions among all subarrays of length k.
//
// Approach: Sliding window with Fenwick tree. Maintain inversion
// count as window moves. Add inversions involving new right element,
// remove inversions involving old left element.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minInversionCount([]int{5, 3, 2, 1, 4}, 4))
	// Example 2
	fmt.Println(minInversionCount([]int{1, 2, 3, 4}, 2))
	// Edge: k=1
	fmt.Println(minInversionCount([]int{5, 3, 1}, 1))
	// Edge: k=n
	fmt.Println(minInversionCount([]int{3, 1, 2}, 3))
}

func minInversionCount(nums []int, k int) int64 {
	n := len(nums)
	if k <= 0 || k > n {
		return 0
	}

	// Coordinate compress
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, n)
	copy(sorted, nums)
	sortInts(sorted)
  // Membuat map untuk pencarian O(1): key → value
	comp := make(map[int]int)
	for i, v := range sorted {
		if i == 0 || v != sorted[i-1] {
			comp[v] = len(comp) + 1
		}
	}

	size := len(comp) + 2
  // Membuat slice untuk menyimpan hasil
	bit := make([]int, size)

	add := func(idx, val int) {
		for idx < size {
			bit[idx] += val
			idx += idx & -idx
		}
	}

	sum := func(idx int) int {
		s := 0
		for idx > 0 {
			s += bit[idx]
			idx -= idx & -idx
		}
		return s
	}

	// First window [0, k-1]
	var inv int64
	for i := 0; i < k; i++ {
		idx := comp[nums[i]]
		// Count existing elements > nums[i]
		inv += int64(sum(size-1) - sum(idx))
		add(idx, 1)
	}

	minInv := inv

	// Slide window
	for i := k; i < n; i++ {
		// Remove leftmost element nums[i-k]
		leftIdx := comp[nums[i-k]]
		// Inversions involving left element = elements after it in window that are smaller
		// When we remove it, we lose these inversions
		invRemoved := int64(sum(leftIdx - 1))
		add(leftIdx, -1)
		inv -= invRemoved

		// Add rightmost element nums[i]
		rightIdx := comp[nums[i]]
		// Inversions involving new element = existing elements > nums[i]
		invAdded := int64(sum(size-1) - sum(rightIdx))
		add(rightIdx, 1)
		inv += invAdded

		if inv < minInv {
			minInv = inv
		}
	}

	return minInv
}

func sortInts(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}
```
