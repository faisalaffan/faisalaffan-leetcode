# 2527 — Find Xor Beauty Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func xorBeauty(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2527: Find Xor-Beauty of Array
// https://leetcode.com/problems/find-xor-beauty-of-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// XOR of all (nums[i]|nums[j]) & nums[k] over all i,j,k = XOR of all nums[i].
// Because the expression simplifies: each bit appears in result iff it appears odd times in nums.

import "fmt"

func main() {
	fmt.Println(xorBeauty([]int{1, 4})) // 5
	fmt.Println(xorBeauty([]int{15, 45, 20, 2, 34, 35, 5, 44, 32, 30})) // 34
}

func xorBeauty(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
```
