# 3314 — Construct The Minimum Bitwise Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ConstructTheMinimumBitwiseArrayI(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * min_val). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #3314: Construct the Minimum Bitwise Array I
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConstructTheMinimumBitwiseArrayI([]int{2, 3, 5, 7}))
	fmt.Println(ConstructTheMinimumBitwiseArrayI([]int{11, 13, 31}))
}

// ConstructTheMinimumBitwiseArrayI returns an array where ans[i] is the smallest number such that ans[i] | (ans[i]+1) == nums[i].
// Time: O(n * min_val). Space: O(n).
func ConstructTheMinimumBitwiseArrayI(nums []int) []int {
  // Alokasi slice
	result := make([]int, len(nums))
	for i, num := range nums {
		found := false
		for candidate := 0; candidate < num; candidate++ {
			if candidate|(candidate+1) == num {
				result[i] = candidate
				found = true
				break
			}
		}
		if !found {
			result[i] = -1
		}
	}
	return result
}
```
