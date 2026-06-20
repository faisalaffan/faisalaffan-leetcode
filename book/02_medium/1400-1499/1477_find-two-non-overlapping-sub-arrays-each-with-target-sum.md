# 1477 — Find Two Non Overlapping Sub Arrays Each With Target Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func minSumOfLengths(arr []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, Prefix Sum

**Waktu:** O(n) where n = len(arr)  |  **Ruang:** O(n) for prefix minimum array

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1477: Find Two Non-overlapping Sub-arrays Each With Target Sum
// https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3)) // 2

	// Test case 2
	fmt.Println(minSumOfLengths([]int{7, 3, 4, 7}, 7)) // 2

	// Test case 3
	fmt.Println(minSumOfLengths([]int{4, 3, 2, 6, 2, 3, 4}, 6)) // -1

	// Test case 4
	fmt.Println(minSumOfLengths([]int{1, 1, 1, 2, 1, 1}, 3)) // 3
}

// Time: O(n) where n = len(arr)
// Space: O(n) for prefix minimum array
func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	// left[i] = minimum length of subarray with sum = target ending at or before i
  // Alokasi slice
	left := make([]int, n)
  // Range loop
	for i := range left {
		left[i] = math.MaxInt32
	}

	prefixSum := 0
  // HashMap: O(1) lookup
	sumMap := make(map[int]int)
	sumMap[0] = -1
	bestLeft := math.MaxInt32

	for i := 0; i < n; i++ {
		prefixSum += arr[i]
		if j, ok := sumMap[prefixSum-target]; ok {
			length := i - j
			if length < bestLeft {
				bestLeft = length
			}
			if i > 0 && left[i-1] < bestLeft {
				bestLeft = left[i-1]
			}
		}
		left[i] = bestLeft
		sumMap[prefixSum] = i
	}

	// right[i] = minimum length of subarray with sum = target starting at or after i
  // Alokasi slice
	right := make([]int, n)
  // Range loop
	for i := range right {
		right[i] = math.MaxInt32
	}

	suffixSum := 0
	sumMap = make(map[int]int)
	sumMap[0] = n
	bestRight := math.MaxInt32

	for i := n - 1; i >= 0; i-- {
		suffixSum += arr[i]
		if j, ok := sumMap[suffixSum-target]; ok {
			length := j - i
			if length < bestRight {
				bestRight = length
			}
			if i < n-1 && right[i+1] < bestRight {
				bestRight = right[i+1]
			}
		}
		right[i] = bestRight
		sumMap[suffixSum] = i
	}

	result := math.MaxInt32
	for i := 0; i < n-1; i++ {
		if left[i] != math.MaxInt32 && right[i+1] != math.MaxInt32 {
			total := left[i] + right[i+1]
			if total < result {
				result = total
			}
		}
	}

	if result == math.MaxInt32 {
		return -1
	}
	return result
}
```
