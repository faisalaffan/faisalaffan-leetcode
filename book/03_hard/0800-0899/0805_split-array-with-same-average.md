# 0805 — Split Array With Same Average

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func splitArraySameAverage(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #805: Split Array With Same Average
// https://leetcode.com/problems/split-array-with-same-average/
// Difficulty: Hard
//
// Given an array A, can we split it into two non-empty subsets B and C such
// that average(B) == average(C)?
//
// A subset with same average as the full set must satisfy:
//   sum(subset) / len(subset) == total / n
//   => sum(subset) == len(subset) * total / n
//
// Approach: Meet-in-the-Middle
//   - Split the array into two halves (n1, n2).
//   - For each half, enumerate all subsets and record achievable sums per size.
//   - A single half can already contain a valid subset.
//   - Otherwise, combine a subset from the left half with one from the right.

import "fmt"

func main() {
	// Example: [1,2,3,4,5,6,7,8] → true (e.g. {1,2,3,6,8} avg=4)
	fmt.Println(splitArraySameAverage([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	// Trivial: [0] → false (need non-empty proper subset)
	fmt.Println(splitArraySameAverage([]int{0}))
	// [1,2,3] → false (avg 2, no proper subset with avg 2)
	fmt.Println(splitArraySameAverage([]int{1, 2, 3}))
	// [3,1] → true ({3} avg=3, {1} avg=1? no... wait {1,3} total=4 avg=2)
	// Actually: {3} avg=3, {1} avg=1. avg != avg, so false.
	fmt.Println(splitArraySameAverage([]int{3, 1}))
	// [0,0,0,0] → true ({0}, {0,0,0} both avg=0)
	fmt.Println(splitArraySameAverage([]int{0, 0, 0, 0}))
}

func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}

	total := 0
	for _, v := range nums {
		total += v
	}

	// Early exit: no k in [1, n-1] satisfies k*total % n == 0
	possible := false
	for k := 1; k < n; k++ {
		if (k*total)%n == 0 {
			possible = true
			break
		}
	}
	if !possible {
		return false
	}

	// Meet-in-the-middle
	n1 := n / 2
	// leftSums[size] = set of achievable sums with given subset size
  // Alokasi slice
	leftSums := make([]map[int]bool, n1+1)
  // Range loop
	for i := range leftSums {
		leftSums[i] = make(map[int]bool)
	}

	for mask := 1; mask < (1 << n1); mask++ {
		sum, size := 0, 0
		for i := 0; i < n1; i++ {
			if mask>>i&1 == 1 {
				sum += nums[i]
				size++
			}
		}
		// Check if this subset alone is a valid split
		if size < n && sum*n == total*size {
			return true
		}
		leftSums[size][sum] = true
	}

	n2 := n - n1
	for mask := 1; mask < (1 << n2); mask++ {
		sum, size := 0, 0
		for i := 0; i < n2; i++ {
			if mask>>i&1 == 1 {
				sum += nums[n1+i]
				size++
			}
		}
		// Check right subset alone
		if size < n && sum*n == total*size {
			return true
		}
		// Combine with left
		for leftSize := 1; leftSize <= n1; leftSize++ {
			totalSize := leftSize + size
			if totalSize <= 0 || totalSize >= n {
				continue
			}
			if (totalSize*total)%n != 0 {
				continue
			}
			targetSum := totalSize * total / n
			needSum := targetSum - sum
			if leftSums[leftSize][needSum] {
				return true
			}
		}
	}

	return false
}
```
