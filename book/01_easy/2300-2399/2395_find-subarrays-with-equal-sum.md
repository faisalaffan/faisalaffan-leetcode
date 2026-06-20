# 2395 — Find Subarrays With Equal Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindSubarraysWithEqualSum(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2395: Find Subarrays With Equal Sum
// https://leetcode.com/problems/find-subarrays-with-equal-sum/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(FindSubarraysWithEqualSum([]int{4, 2, 4}))   // true
	fmt.Println(FindSubarraysWithEqualSum([]int{1, 2, 3, 4, 5})) // false
}

func FindSubarraysWithEqualSum(nums []int) bool {
	seen := map[int]bool{}
	for i := 1; i < len(nums); i++ {
		sum := nums[i-1] + nums[i]
		if seen[sum] {
			return true
		}
		seen[sum] = true
	}
	return false
}
```
