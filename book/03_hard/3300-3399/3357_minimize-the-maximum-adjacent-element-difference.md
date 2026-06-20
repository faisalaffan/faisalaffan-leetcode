# 3357 — Minimize The Maximum Adjacent Element Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minDifference(nums []int) int
```

> **💡 Hint:** Analyze known-adjacent gaps and ranges for missing segments.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3357: Minimize the Maximum Adjacent Element Difference
// https://leetcode.com/problems/minimize-the-maximum-adjacent-element-difference/
// Difficulty: Hard
//
// Given an array nums where some values are missing (denoted by -1), select
// a pair of positive integers (x, y) exactly once and replace each -1 with
// either x or y. Minimize the maximum absolute adjacent difference.
//
// Approach: Analyze known-adjacent gaps and ranges for missing segments.
// Use binary search on answer (max allowed diff) and check feasibility.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(minDifference([]int{1, 2, -1, 10, 8}))
	// Example 2
	fmt.Println(minDifference([]int{-1, -1, -1}))
	// Example 3
	fmt.Println(minDifference([]int{-1, 10, -1, 8}))
	// Edge: no -1
	fmt.Println(minDifference([]int{1, 2, 3}))
	// Edge: single element
	fmt.Println(minDifference([]int{5}))
}

func minDifference(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Find max adjacent diff between known elements
	maxAdj := 0
	for i := 0; i < n-1; i++ {
		if nums[i] > 0 && nums[i+1] > 0 {
			diff := nums[i] - nums[i+1]
			if diff < 0 {
				diff = -diff
			}
			if diff > maxAdj {
				maxAdj = diff
			}
		}
	}

	// Find ranges of known neighbors adjacent to -1 segments
  // Alokasi slice integer
	neighbors := make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			neighbors = append(neighbors, nums[i])
		} else {
			if i > 0 && nums[i-1] > 0 {
				neighbors = append(neighbors, nums[i-1])
			}
			if i+1 < n && nums[i+1] > 0 {
				neighbors = append(neighbors, nums[i+1])
			}
		}
	}

	if len(neighbors) == 0 {
		return 0
	}

	minVal := math.MaxInt32
	maxVal := 0
	for _, v := range neighbors {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	// Binary search on answer
	left, right := maxAdj, maxVal-minVal+maxAdj
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := (left + right) / 2
		if canAchieve(nums, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canAchieve(nums []int, d int) bool {
	n := len(nums)
	// For each segment of -1s, compute the feasible range of values
	// that keeps all adjacent diffs <= d
	type seg struct {
		leftNeighbor  int
		rightNeighbor int
		length        int
	}

	segments := make([]seg, 0)
	i := 0
	for i < n {
		if nums[i] == -1 {
			start := i
			for i < n && nums[i] == -1 {
				i++
			}
			left := -1
			if start > 0 {
				left = nums[start-1]
			}
			right := -1
			if i < n {
				right = nums[i]
			}
			segments = append(segments, seg{left, right, i - start})
		} else {
			i++
		}
	}

	if len(segments) == 0 {
		return true
	}

	// Check if we can choose x and y to satisfy all segments
	// Strategy: try all possible x values from neighbor values
  // Membuat map (HashMap) — pencarian O(1)
	neighborSet := make(map[int]bool)
	for _, v := range nums {
		if v > 0 {
			neighborSet[v] = true
		}
	}

	// Try all pairs from neighbor values (limited set)
  // Alokasi slice integer
	candidates := make([]int, 0, len(neighborSet))
	for v := range neighborSet {
		candidates = append(candidates, v)
	}
	// Also try some computed values
	if len(candidates) > 0 {
		avg := (candidates[0] + candidates[len(candidates)-1]) / 2
		candidates = append(candidates, avg)
	}

	for _, x := range candidates {
		for _, y := range candidates {
			if checkPair(nums, d, x, y) {
				return true
			}
		}
	}

	return false
}

func checkPair(nums []int, d int, x int, y int) bool {
	// Check if using values x and y for -1s satisfies max diff <= d
	n := len(nums)
	prev := -1
	for i := 0; i < n; i++ {
		var cur int
		if nums[i] > 0 {
			cur = nums[i]
		} else {
			// Try both x and y, pick the one that works
			if prev > 0 {
				if abs(cur-prev) > d {
					// Current choice is wrong, use the other
					cur = x + y - cur // switch between x and y
				}
				if abs(cur-prev) > d {
					// Neither works
					return false
				}
			} else {
				cur = x
			}
		}
		prev = cur
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
