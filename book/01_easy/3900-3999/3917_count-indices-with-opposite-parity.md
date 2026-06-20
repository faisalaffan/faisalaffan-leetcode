# 3917 — Count Indices With Opposite Parity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountIndicesWithOppositeParity(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3917: Count Indices With Opposite Parity
// https://leetcode.com/problems/count-indices-with-opposite-parity/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountIndicesWithOppositeParity([]int{1, 2, 3, 4}))
	fmt.Println(CountIndicesWithOppositeParity([]int{2, 4, 6}))
}

// Time: O(n)
// Space: O(1)
func CountIndicesWithOppositeParity(nums []int) []int {
	totalEven, totalOdd := 0, 0
	for _, v := range nums {
		if v%2 == 0 {
			totalEven++
		} else {
			totalOdd++
		}
	}

  // Alokasi slice
	ans := make([]int, len(nums))
	for i, v := range nums {
		if v%2 == 0 {
			totalEven--
			ans[i] = totalOdd
		} else {
			totalOdd--
			ans[i] = totalEven
		}
	}
	return ans
}
```
