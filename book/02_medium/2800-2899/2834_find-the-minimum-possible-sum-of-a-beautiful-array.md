# 2834 — Find The Minimum Possible Sum Of A Beautiful Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindTheMinimumPossibleSumOfABeautifulArray(n int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2834: Find the Minimum Possible Sum of a Beautiful Array
// https://leetcode.com/problems/find-the-minimum-possible-sum-of-a-beautiful-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func FindTheMinimumPossibleSumOfABeautifulArray(n int, target int) int {
  // HashMap: O(1) lookup
	used := make(map[int]bool)
	sum := 0
	for i := 1; len(used) < n; i++ {
		if used[target-i] {
			continue
		}
		used[i] = true
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(2, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(3, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(5, 5))
}
```
