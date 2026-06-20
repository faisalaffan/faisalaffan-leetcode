# 3020 — Find The Maximum Number Of Elements In Subset

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maximumLength(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3020: Find the Maximum Number of Elements in Subset
// https://leetcode.com/problems/find-the-maximum-number-of-elements-in-subset/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maximumLength([]int{5, 4, 1, 2, 2}))
	fmt.Println(maximumLength([]int{1, 3, 2, 4}))
	fmt.Println(maximumLength([]int{1, 1}))
}

func maximumLength(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}

	ans := 0
	seen := map[int]bool{}

	// x = 1 special: 1^2 = 1, all 1's chain together
	if c := cnt[1]; c > 0 {
		seen[1] = true
		if c > ans {
			ans = c
		}
	}

	for x := range cnt {
		if seen[x] || x == 1 {
			continue
		}
		chain := []int{}
		for y := x; cnt[y] > 0 && y <= 1e9; y = y * y {
			chain = append(chain, y)
			seen[y] = true
		}
		if len(chain) == 0 {
			continue
		}
		result := 0
		for i := len(chain) - 1; i >= 0; i-- {
			v := chain[i]
			c := cnt[v]
			if result == 0 {
				result = 1
				c--
			}
			take := c
			if take > 2 {
				take = 2
			}
			result += take
		}
		if result > ans {
			ans = result
		}
	}
	if ans == 0 {
		ans = 1
	}
	return ans
}
```
