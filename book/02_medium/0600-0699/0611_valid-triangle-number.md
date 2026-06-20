# 0611 — Valid Triangle Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func TriangleNumber(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(n^2)  |  **Ruang:** O(log n) for sorting

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #611: Valid Triangle Number
// https://leetcode.com/problems/valid-triangle-number/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(log n) for sorting

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TriangleNumber([]int{2, 2, 3, 4}))
	fmt.Println(TriangleNumber([]int{4, 2, 3, 4}))
}

func TriangleNumber(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	count := 0
	n := len(nums)

	for i := n - 1; i >= 2; i-- {
		left, right := 0, i-1
  // Two-pointer loop
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}
```
