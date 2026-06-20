# 1330 — Reverse Subarray To Maximize Array Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxValueAfterReverse(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1330: Reverse Subarray To Maximize Array Value
// https://leetcode.com/problems/reverse-subarray-to-maximize-array-value/
// Difficulty: Hard
//
// Approach: Math observation + O(n) scan.
// Reversing a subarray [l..r] only changes two boundary terms in the total sum:
//   old: |a[l-1]-a[l]| + |a[r]-a[r+1]|
//   new: |a[l-1]-a[r]| + |a[l]-a[r+1]|
// Using |x| = max(x, -x), expand the delta into 4 separable cases and
// compute the max improvement in O(n). Also handle prefix/suffix reversals.

import "fmt"

func maxValueAfterReverse(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	base := 0
	for i := 0; i < n-1; i++ {
		base += abs(nums[i] - nums[i+1])
	}

	abs := abs
	maxDelta := 0

	// Case A: interior reversal (l >= 1, r <= n-2)
	// Delta = |a[l-1]-a[r]| + |a[l]-a[r+1]| - |a[l-1]-a[l]| - |a[r]-a[r+1]|
	// Expand |x|=max(x,-x) into 4 sign cases, separate into left/right parts.
	for _, signs := range [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		s1, s2 := signs[0], signs[1]
		best := -1 << 30
		for i := 1; i <= n-2; i++ {
			// pair (i-1,i) as left boundary
			left := s1*nums[i-1] + s2*nums[i] - abs(nums[i-1]-nums[i])
			if left > best {
				best = left
			}
			// pair (i,i+1) as right boundary
			right := -s1*nums[i] - s2*nums[i+1] - abs(nums[i]-nums[i+1])
			if best+right > maxDelta {
				maxDelta = best + right
			}
		}
	}

	// Case B: prefix reversal (l = 0)
	// Delta = |a[0]-a[r+1]| - |a[r]-a[r+1]|
	for r := 0; r < n-1; r++ {
		delta := abs(nums[0]-nums[r+1]) - abs(nums[r]-nums[r+1])
		if delta > maxDelta {
			maxDelta = delta
		}
	}

	// Case C: suffix reversal (r = n-1)
	// Delta = |a[l-1]-a[n-1]| - |a[l-1]-a[l]|
	for l := 1; l < n; l++ {
		delta := abs(nums[l-1]-nums[n-1]) - abs(nums[l-1]-nums[l])
		if delta > maxDelta {
			maxDelta = delta
		}
	}

	return base + maxDelta
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example: reverse [3,1,5] to get [2,5,1,3,4], sum=10
	fmt.Println(maxValueAfterReverse([]int{2, 3, 1, 5, 4}))          // 10
	fmt.Println(maxValueAfterReverse([]int{2, 4, 9, 24, 2, 1, 10})) // 68
	fmt.Println(maxValueAfterReverse([]int{1, 2}))                   // 1
	fmt.Println(maxValueAfterReverse([]int{5}))                      // 0
}
```
