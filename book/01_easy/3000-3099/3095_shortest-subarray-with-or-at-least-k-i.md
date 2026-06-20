# 3095 — Shortest Subarray With Or At Least K I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ShortestSubarrayWithOrAtLeastKI(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3095: Shortest Subarray With OR at Least K I
// https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumSubarrayLength
	fmt.Println(ShortestSubarrayWithOrAtLeastKI([]int{1, 2, 3}, 2)) // 1
	fmt.Println(ShortestSubarrayWithOrAtLeastKI([]int{2, 1, 8}, 10)) // 3
	fmt.Println(ShortestSubarrayWithOrAtLeastKI([]int{1, 2}, 10))    // -1
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: minimumSubarrayLength
func ShortestSubarrayWithOrAtLeastKI(nums []int, k int) int {
	n := len(nums)
	minLen := n + 1
	for i := 0; i < n; i++ {
		orVal := 0
		for j := i; j < n; j++ {
			orVal |= nums[j]
			if orVal >= k {
				if j-i+1 < minLen {
					minLen = j - i + 1
				}
				break
			}
		}
	}
	if minLen > n {
		return -1
	}
	return minLen
}
```
