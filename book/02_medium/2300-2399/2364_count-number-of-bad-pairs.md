# 2364 — Count Number Of Bad Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countBadPairs(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2364: Count Number of Bad Pairs
// https://leetcode.com/problems/count-number-of-bad-pairs/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Transform nums[i] into nums[i] - i. Good pairs have same transformed value.
// Bad pairs = total pairs - good pairs.

import "fmt"

func main() {
	fmt.Println(countBadPairs([]int{4, 1, 3, 3})) // 5
	fmt.Println(countBadPairs([]int{1, 2, 3, 4, 5})) // 0
}

func countBadPairs(nums []int) int64 {
	n := len(nums)
  // HashMap: O(1) lookup
	freq := make(map[int]int64)
	var good int64
	for i, v := range nums {
		key := v - i
		good += freq[key]
		freq[key]++
	}
	total := int64(n) * int64(n-1) / 2
	return total - good
}
```
