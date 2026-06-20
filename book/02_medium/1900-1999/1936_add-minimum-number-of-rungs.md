# 1936 — Add Minimum Number Of Rungs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func AddRungs(rungs []int, dist int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1936: Add Minimum Number of Rungs
// https://leetcode.com/problems/add-minimum-number-of-rungs/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(AddRungs([]int{1, 3, 5, 10}, 2))
	fmt.Println(AddRungs([]int{3, 6, 8, 10}, 3))
	fmt.Println(AddRungs([]int{3, 4, 6, 7}, 2))
}

// Time: O(n), Space: O(1)
func AddRungs(rungs []int, dist int) int {
	count := 0
	prev := 0
	for _, r := range rungs {
		gap := r - prev
		if gap > dist {
			count += (gap - 1) / dist
		}
		prev = r
	}
	return count
}
```
