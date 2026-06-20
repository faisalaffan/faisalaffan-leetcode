# 0164 — Maximum Gap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumGap(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n) using bucket sort (Pigeonhole Principle)  |  **Ruang:** O(n) using bucket sort (Pigeonhole Principle)


## 💻 Solusi Go

```go
package main

// LeetCode #164: Maximum Gap
// https://leetcode.com/problems/maximum-gap/
// Difficulty: Medium
// Time: O(n), Space: O(n) using bucket sort (Pigeonhole Principle)

import "fmt"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	if minVal == maxVal {
		return 0
	}

	n := len(nums)
	bucketSize := max(1, (maxVal-minVal)/(n-1))
	bucketCount := (maxVal-minVal)/bucketSize + 1

  // Alokasi slice
	bucketMin := make([]int, bucketCount)
  // Alokasi slice
	bucketMax := make([]int, bucketCount)
  // Range loop
	for i := range bucketMin {
		bucketMin[i] = 1<<31 - 1
		bucketMax[i] = -1 << 31
	}

	for _, num := range nums {
		idx := (num - minVal) / bucketSize
		if num < bucketMin[idx] {
			bucketMin[idx] = num
		}
		if num > bucketMax[idx] {
			bucketMax[idx] = num
		}
	}

	maxGap := 0
	prevMax := minVal
	for i := 0; i < bucketCount; i++ {
		if bucketMin[i] == 1<<31-1 {
			continue
		}
		maxGap = max(maxGap, bucketMin[i]-prevMax)
		prevMax = bucketMax[i]
	}

	return maxGap
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maximumGap([]int{10}))
	fmt.Println(maximumGap([]int{1, 10000000}))
}
```
