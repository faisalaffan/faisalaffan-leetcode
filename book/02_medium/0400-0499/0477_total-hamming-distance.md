# 0477 — Total Hamming Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func TotalHammingDistance(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Bitmask

**Waktu:** O(n * 32) = O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Bitmask** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #477: Total Hamming Distance
// https://leetcode.com/problems/total-hamming-distance/
// Difficulty: Medium
// Time: O(n * 32) = O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(TotalHammingDistance([]int{4, 14, 2}))
	fmt.Println(TotalHammingDistance([]int{4, 14, 4}))
}

func TotalHammingDistance(nums []int) int {
	total := 0
	n := len(nums)

	for bit := 0; bit < 32; bit++ {
		countOnes := 0
		for _, num := range nums {
			if num&(1<<bit) != 0 {
				countOnes++
			}
		}
		countZeros := n - countOnes
		total += countOnes * countZeros
	}

	return total
}
```
