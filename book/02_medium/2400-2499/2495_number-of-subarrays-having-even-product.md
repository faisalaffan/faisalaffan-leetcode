# 2495 — Number Of Subarrays Having Even Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func evenProduct(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2495: Number of Subarrays Having Even Product
// https://leetcode.com/problems/number-of-subarrays-having-even-product/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Product is even if at least one even element.
// Count subarrays from last even position.

import "fmt"

func main() {
	fmt.Println(evenProduct([]int{1, 2, 3})) // 3 ([2], [1,2], [2,3], [1,2,3])
	fmt.Println(evenProduct([]int{1, 3, 5})) // 0
}

func evenProduct(nums []int) int64 {
	var ans int64
	lastEven := -1
	for i, v := range nums {
		if v%2 == 0 {
			lastEven = i
		}
		if lastEven != -1 {
			ans += int64(lastEven + 1)
		}
	}
	return ans
}
```
