# 1099 — Two Sum Less Than K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func twoSumLessThanK(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1099: Two Sum Less Than K
// https://leetcode.com/problems/two-sum-less-than-k/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60)) // 58
	fmt.Println(twoSumLessThanK([]int{10, 20, 30}, 15))                   // -1
}

// LeetCode submission: twoSumLessThanK
func twoSumLessThanK(nums []int, k int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	ans := -1
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum < k {
			if sum > ans {
				ans = sum
			}
			i++
		} else {
			j--
		}
	}
	return ans
}
```
