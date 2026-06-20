# 1365 — How Many Numbers Are Smaller Than The Current Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func smallerNumbersThanCurrent(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1) — since count array is fixed size 101  |  **Ruang:** O(1) — since count array is fixed size 101


## 💻 Solusi Go

```go
package main

// LeetCode #1365: How Many Numbers Are Smaller Than the Current Number
// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
// Difficulty: Easy
//
// LeetCode submission: func smallerNumbersThanCurrent(nums []int) []int

import "fmt"

func main() {
	fmt.Println(HowManyNumbersAreSmallerThanTheCurrentNumber([]int{8, 1, 2, 2, 3})) // [4 0 1 1 3]
	fmt.Println(HowManyNumbersAreSmallerThanTheCurrentNumber([]int{6, 5, 4, 8}))    // [2 1 0 3]
	fmt.Println(HowManyNumbersAreSmallerThanTheCurrentNumber([]int{7, 7, 7, 7}))    // [0 0 0 0]
}

// Time: O(n), Space: O(1) — since count array is fixed size 101
func HowManyNumbersAreSmallerThanTheCurrentNumber(nums []int) []int {
  // Alokasi slice
	count := make([]int, 101)
	for _, v := range nums {
		count[v]++
	}
	for i := 1; i < 101; i++ {
		count[i] += count[i-1]
	}
  // Alokasi slice
	res := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			res[i] = count[v-1]
		}
	}
	return res
}
```
