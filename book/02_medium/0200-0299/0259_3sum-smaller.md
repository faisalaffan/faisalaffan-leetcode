# 0259 — 3Sum Smaller

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func threeSumSmaller(nums []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(n^2), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #259: 3Sum Smaller
// https://leetcode.com/problems/3sum-smaller/
// Difficulty: Medium [Paid]
// Time: O(n^2), Space: O(1)

import (
	"fmt"
	"sort"
)

func threeSumSmaller(nums []int, target int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	count := 0

  // Linear scan O(n)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
  // Two-pointer loop
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < target {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}

func main() {
	fmt.Println(threeSumSmaller([]int{-2, 0, 1, 3}, 2))
	fmt.Println(threeSumSmaller([]int{1, 1, -2}, 1))
	fmt.Println(threeSumSmaller([]int{0, 0, 0}, 0))
}
```
