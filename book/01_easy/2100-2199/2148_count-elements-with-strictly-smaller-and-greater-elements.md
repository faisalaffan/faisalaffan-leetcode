# 2148 — Count Elements With Strictly Smaller And Greater Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountElementsWithStrictlySmallerAndGreaterElements(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2148: Count Elements With Strictly Smaller and Greater Elements
// https://leetcode.com/problems/count-elements-with-strictly-smaller-and-greater-elements/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountElementsWithStrictlySmallerAndGreaterElements([]int{11, 7, 2, 15}))   // 2
	fmt.Println(CountElementsWithStrictlySmallerAndGreaterElements([]int{-3, 3, 3, 90}))   // 2
	fmt.Println(CountElementsWithStrictlySmallerAndGreaterElements([]int{1, 2, 3}))         // 1
}

// Time: O(n), Space: O(1)
func CountElementsWithStrictlySmallerAndGreaterElements(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if min == max {
		return 0
	}

	count := 0
	for _, v := range nums {
		if v > min && v < max {
			count++
		}
	}
	return count
}
```
