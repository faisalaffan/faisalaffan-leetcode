# 2090 — K Radius Subarray Averages

## Deskripsi

**Soal:** [2090. K Radius Subarray Averages](https://leetcode.com/problems/k-radius-subarray-averages/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func getAverages(nums []int, k int) []int`

## Solusi Go

```go
package main

// LeetCode #2090: K Radius Subarray Averages
// https://leetcode.com/problems/k-radius-subarray-averages/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func getAverages(nums []int, k int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
  // Iterasi seluruh elemen
	for i := range result {
		result[i] = -1
	}

	if n < 2*k+1 {
		return result
	}

	// Prefix sum
  // Membuat slice untuk menyimpan hasil
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
