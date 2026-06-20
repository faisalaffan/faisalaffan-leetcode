# 3471 — Find The Largest Almost Missing Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindTheLargestAlmostMissingInteger(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3471: Find the Largest Almost Missing Integer
// https://leetcode.com/problems/find-the-largest-almost-missing-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheLargestAlmostMissingInteger([]int{3, 9, 2, 3, 1, 6, 7, 8, 9}, 2))
	fmt.Println(FindTheLargestAlmostMissingInteger([]int{0, 0}, 1))
}

// FindTheLargestAlmostMissingInteger returns the largest integer that appears fewer than k times in nums.
// Time: O(n). Space: O(n).
func FindTheLargestAlmostMissingInteger(nums []int, k int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	largest := -1
	for val, count := range freq {
		if count < k && val > largest {
			largest = val
		}
	}
	return largest
}
```
