# 0667 — Beautiful Arrangement Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func constructArray(n int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #667: Beautiful Arrangement II
// https://leetcode.com/problems/beautiful-arrangement-ii/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(constructArray(3, 1))
	fmt.Println(constructArray(3, 2))
}

func constructArray(n int, k int) []int {
  // Alokasi slice
	result := make([]int, n)
	left, right := 1, k+1

	for i := 0; i <= k; i++ {
		if i%2 == 0 {
			result[i] = left
			left++
		} else {
			result[i] = right
			right--
		}
	}

	for i := k + 1; i < n; i++ {
		result[i] = i + 1
	}

	return result
}
```
