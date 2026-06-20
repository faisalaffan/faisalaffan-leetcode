# 3912 — Valid Elements In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func ValidElementsInAnArray(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3912: Valid Elements in an Array
// https://leetcode.com/problems/valid-elements-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ValidElementsInAnArray([]int{1, 2, 4, 2, 3, 2}))
	fmt.Println(ValidElementsInAnArray([]int{5, 5, 5, 5}))
	fmt.Println(ValidElementsInAnArray([]int{1}))
}

// Time: O(n)
// Space: O(n)
func ValidElementsInAnArray(nums []int) []int {
	n := len(nums)
	isGreaterRight := make([]bool, n)

	maxRight := -1
	for i := n - 1; i >= 0; i-- {
		if nums[i] > maxRight {
			isGreaterRight[i] = true
			maxRight = nums[i]
		}
	}

	var result []int
	maxLeft := -1
	for i := 0; i < n; i++ {
		if nums[i] > maxLeft || isGreaterRight[i] {
			result = append(result, nums[i])
		}
		if nums[i] > maxLeft {
			maxLeft = nums[i]
		}
	}
	return result
}
```
