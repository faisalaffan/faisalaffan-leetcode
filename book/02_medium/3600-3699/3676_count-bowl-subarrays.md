# 3676 — Count Bowl Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countBowlSubarrays(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3676: Count Bowl Subarrays
// https://leetcode.com/problems/count-bowl-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countBowlSubarrays(nums []int) int64 {
	var ans int64 = 0
	var stack []int

	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			if len(stack) >= 2 {
				ans++
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	return ans
}

func main() {
	fmt.Println(countBowlSubarrays([]int{1, 3, 5, 4, 2}))
	fmt.Println(countBowlSubarrays([]int{3, 1, 2, 4}))
	fmt.Println(countBowlSubarrays([]int{1, 2, 3, 4}))
}
```
