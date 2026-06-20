# 3469 — Find Minimum Cost To Remove Array Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minCost(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3469: Find Minimum Cost to Remove Array Elements
// https://leetcode.com/problems/find-minimum-cost-to-remove-array-elements/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func minCost(nums []int) int {
	n := len(nums)
  // Alokasi slice
	f := make([]int, n)
	if n%2 == 0 {
		for i, x := range nums {
			f[i] = max(x, nums[n-1])
		}
	} else {
		copy(f, nums)
	}
	for i := n - 3 + n%2; i > 0; i -= 2 {
		b, c := nums[i], nums[i+1]
		for j := 0; j < i; j++ {
			a := nums[j]
			f[j] = min(
				f[j]+max(b, c),
				f[i]+max(a, c),
				f[i+1]+max(a, b),
			)
		}
	}
	return f[0]
}

func main() {
	fmt.Println(minCost([]int{6, 2, 8, 4})) // 12
	fmt.Println(minCost([]int{1, 2, 3}))     // 3
	fmt.Println(minCost([]int{5}))            // 5
}
```
