# 1989 — Maximum Number Of People That Can Be Caught In Tag

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MaximumNumberOfPeopleThatCanBeCaughtInTag(team []int, dist int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1989: Maximum Number of People That Can Be Caught in Tag
// https://leetcode.com/problems/maximum-number-of-people-that-can-be-caught-in-tag/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(MaximumNumberOfPeopleThatCanBeCaughtInTag([]int{0, 1, 1, 0}, 2))
	fmt.Println(MaximumNumberOfPeopleThatCanBeCaughtInTag([]int{1, 1, 0, 0}, 1))
	fmt.Println(MaximumNumberOfPeopleThatCanBeCaughtInTag([]int{0, 0, 1, 0, 1, 0, 0, 1, 1, 0}, 5))
}

// Time: O(n), Space: O(1)
func MaximumNumberOfPeopleThatCanBeCaughtInTag(team []int, dist int) int {
	ans := 0
	n := len(team)
	j := 0
	for i := 0; i < n; i++ {
		if team[i] == 1 {
			for j < n && (team[j] == 1 || i-j > dist) {
				j++
			}
			if j < n && abs1989(i-j) <= dist {
				ans++
				j++
			}
		}
	}
	return ans
}

func abs1989(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
