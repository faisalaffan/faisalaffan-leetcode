# 1099 — Two Sum Less Than K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func twoSumLessThanK(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

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
  // Sort O(n log n)
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
