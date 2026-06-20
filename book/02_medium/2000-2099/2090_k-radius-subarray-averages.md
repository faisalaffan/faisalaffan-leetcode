# 2090 — K Radius Subarray Averages

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func getAverages(nums []int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2090: K Radius Subarray Averages
// https://leetcode.com/problems/k-radius-subarray-averages/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func getAverages(nums []int, k int) []int {
	n := len(nums)
  // Alokasi slice
	result := make([]int, n)
  // Range loop
	for i := range result {
		result[i] = -1
	}

	if n < 2*k+1 {
		return result
	}

	// Prefix sum
  // Alokasi slice
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	for i := k; i < n-k; i++ {
		sum := prefix[i+k+1] - prefix[i-k]
		result[i] = int(sum / int64(2*k+1))
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
	// Expected: [-1,-1,-1,5,4,4,-1,-1,-1]

	// Test case 2
	fmt.Println("Test 2:", getAverages([]int{100000}, 0))
	// Expected: [100000]

	// Test case 3
	fmt.Println("Test 3:", getAverages([]int{8}, 100000))
	// Expected: [-1]
}
```
