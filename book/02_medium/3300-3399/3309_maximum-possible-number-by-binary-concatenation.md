# 3309 — Maximum Possible Number By Binary Concatenation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maxGoodNumber(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1) Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3309: Maximum Possible Number by Binary Concatenation
// https://leetcode.com/problems/maximum-possible-number-by-binary-concatenation/
// Difficulty: Medium
// Time: O(1) Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxGoodNumber([]int{1, 2, 3}))    // 30
	fmt.Println(maxGoodNumber([]int{2, 8, 16}))   // 1296
	fmt.Println(maxGoodNumber([]int{1, 1, 1}))    // 7
}

func maxGoodNumber(nums []int) int {
	// Try all 6 permutations
	perms := [][]int{
		{nums[0], nums[1], nums[2]},
		{nums[0], nums[2], nums[1]},
		{nums[1], nums[0], nums[2]},
		{nums[1], nums[2], nums[0]},
		{nums[2], nums[0], nums[1]},
		{nums[2], nums[1], nums[0]},
	}

	maxVal := 0
	for _, p := range perms {
		val := 0
		for _, x := range p {
			bits := 0
			temp := x
			for temp > 0 {
				bits++
				temp >>= 1
			}
			if x == 0 {
				bits = 1
			}
			val = (val << bits) | x
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
```
